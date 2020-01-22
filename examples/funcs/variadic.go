package main

import "fmt"

func whatever(vals ...interface{}) {
	fmt.Printf("%T, %d\n", vals, len(vals))

	for i, val := range vals {
		fmt.Printf("%d, %#v\n", i, val)
	}
}

func sum(numbers ...int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

func main() {
	fmt.Println(sum(10, 20, 30, 40, 50))

	whatever("hi", true, 42, nil)
}
