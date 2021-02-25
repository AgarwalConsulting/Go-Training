package main

import "fmt"

func main() {
	var a []string // slice of strings

	fmt.Println(a, len(a), cap(a))

	if a == nil {
		fmt.Println("Slices are initialized to nil!")
	}

	b := []string{} // Slice literal syntax

	fmt.Println(b, len(b), cap(b))

	if b == nil {
		fmt.Println("Empty slice is nil!")
	}
}
