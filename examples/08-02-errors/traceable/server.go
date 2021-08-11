package main

// Matryoshka: https://www.therussianstore.com/media/wysiwyg/Traditional_Russian_Matryoshka.jpg

import (
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
	connectErr := &DatabaseError{time.Now(), "unable to connect"}

	networkErr := NetworkError("unable to reach server")
	timeoutErr := NetworkError("timeout on receiving data from server")

	errs := []error{networkErr, timeoutErr, dbErr, connectErr, err}

	return errs[(1+rand.Int())%len(errs)] // Returns one of the errors from the `errs` slice
}

func fetchData() error {
	err := randomError()

	// return err

	return NewTraceableError(debug.Stack(), err)
}

func init() {
	fmt.Println("Initializing...")
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Println("Executing Main...")
	var err error // <nil>

	fmt.Printf("\t%T %v\n", err, err) // <nil>, <nil>

	if err != nil {
		err = fetchData()

		fmt.Printf("\t%T %v\n", err, err) // TraceableError, <stack-trace>

		// Handle different errors

		// If underlying error is "unable to reach server"; alert SRE

		// If underlying error is "DB: permission error"; alert DBA
	}
}
