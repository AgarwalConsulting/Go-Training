package main

import "fmt"

func main() {
	var a []string // slice of strings

	fmt.Println(a)

	if a == nil {
		fmt.Println("Slices are initialized to nil!")
	}

	a = []string{}

	if a == nil {
		fmt.Println("Empty slice is nil!")
	}
}
