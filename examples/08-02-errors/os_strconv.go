package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("123ax")

	if err != nil {
		fmt.Printf("%T %v\n", err, err)
	}

	fmt.Println("Num:", i)

	f, err := os.Open("/tmp/unicorn_don't_exist")
	defer f.Close()

	if err != nil {
		fmt.Printf("%T %v\n", err, err)
	}

	err = errors.New("some failure")
	if err != nil {
		fmt.Printf("%T %v\n", err, err)
	}
}
