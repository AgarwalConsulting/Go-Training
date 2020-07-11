package main

import (
	"fmt"
	"strconv"
)

func stringToIntExample() {
	var k int
	var f = "42"

	k, _ = strconv.Atoi(f)

	fmt.Printf("%v, %T\n", f, f)
	fmt.Printf("%v, %T\n", k, k)
}

func integerToString() {
	var i = 42
	var s string

	// s = string(i)
	s = strconv.Itoa(i)

	fmt.Printf("%v, %T\n", s, s)
}

func main() {
	integerToString()
}
