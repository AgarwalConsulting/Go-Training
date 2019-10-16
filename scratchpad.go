package main

import "fmt"

// Add adds two integers
func Add(x, y int) int {
	return x + y
}

func main() {
	var i, j = 42, 50

	fmt.Println("Hello World! Result: ", Add(i, j))

	fmt.Println(&i, &j)
}
