package main

import "fmt"

type myInt int

// type Adder func(int) int

// func(int)
// func fn(x int) {

// }

// func(int) int
// func adderFn(x int) int {
// 	return x + 2
// }

func newAdder(n int) func(int) int {
	return func(x int) int {
		return x + n
	}
}

func main() {
	addBy2 := newAdder(2)

	fmt.Printf("%T\n", addBy2)

	fmt.Println(addBy2(10))
}
