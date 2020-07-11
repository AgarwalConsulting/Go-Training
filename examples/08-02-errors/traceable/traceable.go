package main

// Matryoshka: https://www.therussianstore.com/media/wysiwyg/Traditional_Russian_Matryoshka.jpg

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"time"
)

func randomError() error {
	// fileOpeningErr := *os.PathError{Op: "file not found"}

	_, err := os.Open("/tmp/unicorn.go")

	dbErr := &DatabaseError{time.Now(), "permission error"}
	networkErr := NetworkError("unable to reach server")

	errs := []error{networkErr, dbErr, err}

	return errs[(1+rand.Int())%3]
}

func fetchData() error {
	err := randomError()

	// return err

	return NewTraceableError(debug.Stack(), err)
}

func main() {
	var err error
	if err = fetchData(); err != nil {
		fmt.Println(err)
	}

	var n = NetworkError("unable to reach server")
	// Strict equality comparison iteratively on each of the errors in the chain
	if errors.Is(err, n) {
		fmt.Println("Server is down!")
	}

	var e *DatabaseError
	// If any of the errors in the chain can be type asserted to given variable
	if errors.As(err, &e); e != nil {
		fmt.Println("Database Error: ", e)
	}
}
