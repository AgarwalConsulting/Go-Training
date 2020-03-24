package main

import "fmt"

func main() {
	var a [2]string

	fmt.Printf("a is of type: %T\n", a)

	b := []string{}

	fmt.Printf("b is of type: %T\n", b)
}
