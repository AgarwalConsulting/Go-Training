package main

import "fmt"

func Adder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	var addBy2 func(int) int

	fmt.Println(addBy2)

	addBy2 = Adder(2)

	fmt.Printf("%T\n", addBy2)

	fmt.Println(addBy2(10))
}
