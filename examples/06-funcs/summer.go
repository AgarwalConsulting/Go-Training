package main

import "fmt"

// go:noinline
func summer() func(int) int {
	var sum int

	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum1 := summer()
	sum2 := summer()

	fmt.Println(sum1(10)) //
	fmt.Println(sum1(20)) //
	fmt.Println(sum1(30)) //

	fmt.Println(sum2(12)) //
	fmt.Println(sum2(20)) //
	fmt.Println(sum2(30)) //
}
