package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	// errorMsg := fmt.Sprintf("at %v, %s", time.Now(), "it didn't work")
	// return errors.New(errorMsg)

	// return fmt.Errorf("at %v, %s", time.Now(), "it didn't work")

	return &MyError{time.Now(), "it didn't work"}
}

func main() {
	var err error
	if err = run(); err != nil {
		fmt.Println(err)
	}
}
