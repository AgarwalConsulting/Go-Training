package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type CustomError struct {
	What string
	When time.Time
}

func (e CustomError) Error() string {
	return e.What
}

func (e CustomError) WhenDidItHappen() time.Time {
	return e.When
}

// func doSomething() *CustomError {
// 	var e *CustomError // <nil>

// 	// if rand.Int()%2 == 0 {
// 	// 	c := CustomError("Oops! something bad happened")

// 	// 	e = &c
// 	// }

// 	return e
// }

func doSomething() error {
	if rand.Int()%2 == 0 {
		c := CustomError{"Oops! something bad happened", time.Now()}

		return &c
	}

	return nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var err error

	err = doSomething()

	if err != nil { // Works as expected with Concrete types
		fmt.Println("Error occurred:", err)

		c := err.(*CustomError)
		fmt.Println("Occurred At:", c.When)
		os.Exit(1)
	}

	fmt.Println("Successfully completed!")
}
