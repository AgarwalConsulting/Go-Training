package main

import "fmt"

func Adder(n int) func(int) int {
	return func(x int) int { // x = 10; n = 2
		return x + n // Property of closures: Capture outside scope
	}
}

func main() {
	var addBy2 func(int) int

	// f := Adder

	// fmt.Printf("%T\n", f) // func(int) func(int) int

	// addBy2 = Adder // ?

	fmt.Println(addBy2)

	addBy2 = Adder(2)

	fmt.Printf("%T\n", addBy2)

	fmt.Println(addBy2(10)) //
}
