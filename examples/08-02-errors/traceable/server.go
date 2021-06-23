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
	var err error // Value: <nil>

	fmt.Printf("\t%T %v\n", err, err) // Type: <nil>; Value: <nil>

	err = fetchData()

	fmt.Printf("\t%T\n", err) // Type: main.TraceableError

	if err != nil {
		// fmt.Println("Error encountered:", err)

		ne := NetworkError("unable to reach server")

		var te TraceableError
		te = err.(TraceableError) // Type assertion
		underlyingErr := te.Unwrap()

		fmt.Printf("Underlying error: %T | %v\n", underlyingErr, underlyingErr)

		// if underlyingErr == ne { //
		if errors.Is(err, ne) {
			fmt.Println("Alerting the SRE...")
			fmt.Println("Server is down. Please try again later...")
		}

		var de *DatabaseError

		if errors.As(err, &de) {
			fmt.Println("Alerting DBA...")
		}
	}
}
