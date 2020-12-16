package main

import (
	"fmt"
)

func hello() (int, string) {
	return 42, "Hello"
}

func main() {
	var i int
	// var j string
	i, j := hello()

	fmt.Printf("%T, %v\n", i, i)
	fmt.Printf("%T, %v\n", j, j)
}
