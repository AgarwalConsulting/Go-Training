package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fn1 := adder()
	fn2 := adder()

	fmt.Printf("%T, %#v\n", fn1, fn1)

	fmt.Println(fn1(10)) // 10
	fmt.Println(fn1(4))  // 14

	fmt.Println(fn2(5)) // 5
}
