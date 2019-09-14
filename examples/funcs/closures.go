package main

import "fmt"

// func add(numbers ...int) int {
// 	sum := 0

// 	for _, num := range numbers {
// 		sum += num
// 	}

// 	return sum
// 	// _ = sum * 10

// 	// fmt.Println(mul)
// }

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// sum := add(10, 4, 5)

	// fmt.Println(sum * 10)
	fn1 := adder()
	fn2 := adder()

	fmt.Printf("%T, %#v\n", fn1, fn1)

	fmt.Println(fn1(10)) // 10
	fmt.Println(fn1(4))  // 4 or 14 or 10 ? ->

	fmt.Println(fn2(5)) // 5
}
