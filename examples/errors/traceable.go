package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime/debug"
	"time"
)

type DatabaseError struct {
	When  time.Time
	Cause string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("at %v had %v", e.When, e.Cause)
}

func randomError() error {
	dbErr := &DatabaseError{time.Now(), "permission error"}
	networkErr := NetworkError("unable to reach server")

	errs := []error{dbErr, networkErr}

	return errs[(1+rand.Int())%2]
}

type TraceableError struct {
	Stack []byte
	error
}

func NewTraceableError(err error) TraceableError {
	return TraceableError{debug.Stack(), err}
}

func (t TraceableError) Error() string {
	return fmt.Sprintf("%v\n  At: %s", t.error, t.Stack)
}

func run() error {
	err := randomError()

	return NewTraceableError(err)
}

type NetworkError string

func (e NetworkError) Error() string {
	return string(e)
}

func main() {
	var err error
	if err = run(); err != nil {
		fmt.Println(err)
	}

	var e *DatabaseError
	if errors.As(err, &e); e != nil {
		fmt.Println("Database Error: ", e)
	}
}
