package main

import (
	"database/sql"
	"fmt"
	"log"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type Person struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("sqlite3", "file:/tmp/person.sqlite")
	// db, err := sqlx.Open("postgres", "postgres://localhost/template1?sslmode=disable")
	defer db.Close()

	if err != nil {
		log.Fatal("Unable to connect:", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, name TEXT, age INTEGER);")

	if err != nil {
		log.Fatal("unable to create table: ", err)
	}

	rows, err := db.Query("SELECT * FROM people WHERE name = $1", "Gaurav")

	if err != nil {
		log.Fatal("unable to fetch rows:", err)
	}

	if rows.Next() {
		var (
			id, age int
			name    string
		)
		err := rows.Scan(&id, &name, &age)

		if err != nil {
			log.Fatal("unable to scan the next row:", err)
		}

		p := Person{id, name, age}

		fmt.Println("Next Row: ", p)
	}

	insertStmt, err := db.Prepare("INSERT INTO people(id, name, age) VALUES( $1, $2, $3 )")
	if err != nil {
		log.Fatal("unable to create prepared statement:", err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(1, "Gaurav", 29)
	if err != nil {
		log.Fatal("unable to insert:", err)
	}
}
