package main

import "fmt"

type Any interface {
}

func main() {
	var a interface{}

	a = 42
	describe(a)

	a = false
	describe(a)

	a = "Hello"
	describe(a)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
