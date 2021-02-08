package main

import "fmt"

func main() {
	// var (
	// 	i string
	// 	j int
	// )
	// i, j = "hello", 42

	var i, j = "hello", 42 // Types are inferred

	fmt.Printf("%T\n", i) // string
	fmt.Printf("%T\n", j) // int

	// i = j

	fmt.Println(i, j)
}
