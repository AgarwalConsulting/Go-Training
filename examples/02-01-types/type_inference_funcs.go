package main

import "fmt"

func hello() (string, int) {
	return "hello", 42
}

func main() {
	// var (
	// 	i string
	// 	j int
	// )

	var i, j = hello()

	fmt.Printf("%T\n", i) // string
	fmt.Printf("%T\n", j) // int

	fmt.Println(i, j)
}
