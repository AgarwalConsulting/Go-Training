package main

import "fmt"

func adder() func(int) int {
	var sum = 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	var fn func(int) int

	fmt.Println(fn) // nil

	// fmt.Println(fn(42)) // ?

	fn1 := adder() // func(int) int
	fn2 := adder() // func(int) int

	fmt.Printf("%T, %#v\n", fn1, fn1)

	fmt.Println(fn1(10)) // 10
	fmt.Println(fn1(4))  // 14

	fmt.Println(fn2(5)) // 5
}
