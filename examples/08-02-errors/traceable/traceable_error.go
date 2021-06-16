package main

import (
	"fmt"
)

type TraceableError struct {
	Stack []byte
	// error
	error error
}

func (t TraceableError) Error() string {
	return fmt.Sprintf("%v\n  At: %s", t.error, t.Stack)
}

// Unwrap is used by errors.Is & errors.As
func (t TraceableError) Unwrap() error {
	return t.error
}

func NewTraceableError(stack []byte, err error) TraceableError {
	return TraceableError{stack, err}
}
