package main

import "fmt"

func sum(elements ...interface{}) int {
	sum := 0
	for _, number := range elements {
		if num, ok := number.(int); ok {
			sum += num
		}
	}
	return sum
}

func main() {
	// a := []{10, 20, 30}

	// a = append(a, 40, 50, 60)

	// fmt.Println(a)

	// fmt.Println(sum(a...))

	fmt.Println(sum(10, 20.12, "hello"))
}
