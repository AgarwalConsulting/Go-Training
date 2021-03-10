package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type CustomError string

func (e CustomError) Error() string {
	return string(e)
}

// func doSomething() *CustomError {
// 	var e *CustomError // <nil>

// 	// if rand.Int()%2 == 0 {
// 	// 	c := CustomError("Oops! something bad happened")

// 	// 	e = &c
// 	// }

// 	return e
// }

func doSomething() *CustomError {
	if rand.Int()%2 == 0 {
		c := CustomError("Oops! something bad happened")

		return &c
	}

	return nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// var err *CustomError

	// var err error
	// e1 := doSomething()
	var e1 *CustomError
	e1 = doSomething()

	// fmt.Printf("%T %v\n", err, err) // *CustomError, nil

	if e1 != nil { // Works as expected with Concrete types
		fmt.Println("Error occurred:", e1)
		os.Exit(1)
	}

	fmt.Println("Successfully completed!")
}
