package main

import "fmt"

func main() {
	var i, j = 0.1, 0.2

	fmt.Printf("%T %v\n", i, i)
	fmt.Printf("%T %v\n", j, j)

	fmt.Println(i + j) // 0.3
}
