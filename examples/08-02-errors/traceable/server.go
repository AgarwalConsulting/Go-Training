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

	return errs[(1+rand.Int())%len(errs)]
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
	var err error
	err = fetchData()

	if err != nil {
		fmt.Printf("%T\n", err) // Type: main.TraceableError
		underlyingErr := err.(TraceableError).Unwrap()
		fmt.Printf("\tUnderlying error type: %T\n", underlyingErr)
		fmt.Println(underlyingErr)

		ne := NetworkError("unable to reach server")
		if errors.Is(err, ne) {
			fmt.Println("\t | Server is down. Try again later...")
			fmt.Println("\t | Alerting SRE...")
		}

		var de *DatabaseError
		// de := &DatabaseError{time.Now(), "permission error"}
		if errors.As(err, &de) {
			fmt.Println("\t | Alerting DBA...")
			fmt.Println("\t |", de)
		}
	}
}
