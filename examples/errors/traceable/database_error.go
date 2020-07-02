package main

import (
	"fmt"
	"time"
)

type DatabaseError struct {
	When  time.Time
	Cause string
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("at %v had %v", e.When, e.Cause)
}
