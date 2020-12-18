package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var err error

	i, err := strconv.Atoi("123ax")

	if err != nil {
		fmt.Printf("%T %v\n", err, err)
	}

	fmt.Println("Num:", i)

	// ---

	f, err := os.Open("/tmp/people.json")
	defer f.Close()

	if err != nil {
		fmt.Printf("%T %v\n", err, err)
		os.Exit(1)
	} else {
		fmt.Println("Found the file!")
	}

	// ---

	err = errors.New("some failure")
	if err != nil {
		fmt.Printf("%T %v\n", err, err)
	}

	// ---

	fmt.Println("Program will continue!")
}
