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
	// var err error
	// err = doSomething()

	err := doSomething()

	if err != nil { // Works as expected with Concrete types
		fmt.Println("Error occurred:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully completed!")
}
