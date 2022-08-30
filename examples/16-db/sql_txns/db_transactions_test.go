package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type Queryable interface {
	Query(string, ...any) (*sql.Rows, error)
}

func getCurrentAmount(db Queryable, id int) (int, error) {
	var currentAmount int

	rows, err := db.Query("SELECT amount FROM txns WHERE id = $1;", id)

	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, errors.New("no entry found!")
	}

	rows.Scan(&currentAmount)
	return currentAmount, nil
}

func performTxnUpdateWithRetry(db *sql.DB, t *testing.T, id int, op func(int) int, retryCount int) {
	isRetrying := false

	msg := fmt.Sprintf("Updating id: %d", id)

	if retryCount == 0 {
		log.Error("Abandoning op on:", msg)
		return
	}

	ctx, _ := context.WithCancel(context.Background())

	// tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	assert.Nil(t, err, msg)
	defer func() {
		if !isRetrying {
			tx.Commit()
		}
	}()

	currentAmount, err := getCurrentAmount(tx, id)

	assert.Nil(t, err, msg)

	log.Info(msg, " ", op(currentAmount))

	res, err := tx.Exec("UPDATE txns SET amount = $1 WHERE id = $2;", op(currentAmount), id)

	if err == nil {
		rowsAffected, err := res.RowsAffected()
		assert.Nil(t, err)
		assert.Equal(t, int64(1), rowsAffected)
	} else {
		log.Infof("Retrying %s due to: %v\n", msg, err)
		isRetrying = true
		err := tx.Rollback()

		assert.Nil(t, err)

		performTxnUpdateWithRetry(db, t, id, op, retryCount-1)
	}
}

func performTxnUpdate(db *sql.DB, t *testing.T, id int, op func(int) int, done func()) {
	defer done()

	performTxnUpdateWithRetry(db, t, id, op, 5)
}

func performTxnUpdateInplaceWithRetry(db *sql.DB, t *testing.T, id int, increment bool, retryCount int) {
	msg := fmt.Sprintf("Updating id: %d", id)

	isRetrying := false
	if retryCount == 0 {
		log.Error("Abandoning op on:", msg)
		return
	}

	ctx, _ := context.WithCancel(context.Background())

	// tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	tx, err := db.BeginTx(ctx, &sql.TxOptions{}) // Default: Read Committed
	assert.Nil(t, err, msg)
	defer func() {
		if !isRetrying {
			tx.Commit()
		}
	}()

	// currentAmount, err := getCurrentAmount(tx, id)

	// assert.Nil(t, err, msg)

	// log.Info(msg, " ", op(currentAmount))

	var query string

	if increment {
		query = "UPDATE txns SET amount = (amount + 1) WHERE id = $1;"
	} else {
		query = "UPDATE txns SET amount = (amount - 1) WHERE id = $1;"
	}

	res, err := tx.Exec(query, id)

	// assert.Nil(t, err, msg)

	if err == nil {
		rowsAffected, err := res.RowsAffected()
		assert.Nil(t, err)
		assert.Equal(t, int64(1), rowsAffected)
	} else {
		log.Infof("Retrying %s due to: %v\n", msg, err)
		isRetrying = true
		err := tx.Rollback()

		assert.Nil(t, err)

		time.Sleep(100 * time.Duration(6-retryCount) * time.Millisecond)
		performTxnUpdateInplaceWithRetry(db, t, id, increment, retryCount-1)
		// log.Infof("Underlying Error: %T | %#v\n", err, err)
	}
}

// performTxnUpdateInplace doesn't read current amount, rather it increments or decrements by... without any retry! No SELECT, Only UPDATE
func performTxnUpdateInplace(db *sql.DB, t *testing.T, id int, increment bool, done func()) {
	defer done()

	performTxnUpdateInplaceWithRetry(db, t, id, increment, 5)
}

func increment(x int) int { return x + 1 }
func decrement(x int) int { return x - 1 }

func idToName(id int) string {
	return fmt.Sprintf("%c", 64+id)
}

func TestDBTxnConsistency(t *testing.T) {
	initialAmount := 100

	defer log.Info("Terminating gracefully...")
	db, err := sql.Open("postgres", "postgres://localhost:5432/txn-check?sslmode=disable")

	assert.Nil(t, err, "unable to open db conn")
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS txns (id numeric primary key, name text, amount numeric);")

	assert.Nil(t, err, "unable to create table")

	_, err = db.Exec("DELETE FROM txns")

	assert.Nil(t, err, "unable to delete entries")

	log.Println("Preparing statement...")
	stmt, err := db.Prepare("INSERT INTO txns (id, name, amount) VALUES($1, $2, $3);")

	assert.Nil(t, err, "unable to prepare stmt")

	defer stmt.Close()

	for i := 0; i < 5; i++ {
		id := i + 1

		res, err := stmt.Exec(id, idToName(id), initialAmount)

		if err != nil {
			log.Println("Unable to insert:", err)
			return
		}

		rowsAffected, err := res.RowsAffected()
		assert.Nil(t, err)
		assert.Equal(t, int64(1), rowsAffected)
	}

	// Pseudocode
	// For the first 2 create 20 debits of 1 amount each
	// For the next 2 create 20 credits of 1 amount each
	// For the last one create 20 debits & credits of 1 amount each at the same time
	// Assert the values

	iterCount := 20

	expectedAmounts := map[int]int{
		1: initialAmount - iterCount,
		2: initialAmount - iterCount,
		3: initialAmount + iterCount,
		4: initialAmount + iterCount,
		5: initialAmount,
	}

	var wg sync.WaitGroup
	for i := 0; i < iterCount; i++ {
		wg.Add(6)
		go performTxnUpdate(db, t, 1, decrement, wg.Done)
		// go performTxnUpdateInplace(db, t, 1, false, wg.Done) // UPDATE Only (Increment by & Decrement by)
		go performTxnUpdate(db, t, 2, decrement, wg.Done)
		// go performTxnUpdateInplace(db, t, 2, false, wg.Done) // UPDATE Only (Increment by & Decrement by)

		go performTxnUpdate(db, t, 3, increment, wg.Done)
		// go performTxnUpdateInplace(db, t, 3, true, wg.Done) // UPDATE Only (Increment by & Decrement by)
		go performTxnUpdate(db, t, 4, increment, wg.Done)
		// go performTxnUpdateInplace(db, t, 4, true, wg.Done) // UPDATE Only (Increment by & Decrement by)

		go performTxnUpdate(db, t, 5, increment, wg.Done)
		// go performTxnUpdateInplace(db, t, 5, true, wg.Done) // UPDATE Only (Increment by & Decrement by)
		go performTxnUpdate(db, t, 5, decrement, wg.Done)
		// go performTxnUpdateInplace(db, t, 5, false, wg.Done) // UPDATE Only (Increment by & Decrement by)
	}

	wg.Wait()

	for id, expectedAmount := range expectedAmounts {
		errMsg := fmt.Sprintf("Amounts for %s", idToName(id))

		actualAmount, err := getCurrentAmount(db, id)

		assert.Nil(t, err, errMsg)

		assert.Equal(t, expectedAmount, actualAmount, errMsg)
	}
}
