package main

import (
	"fmt"
)

func describe(a Abser) {
	fmt.Printf("%T %v\n", a, a) // *Vertex, nil

	if a == nil {
		fmt.Println("\tValue is nil?")
	}

	fmt.Println(a.Abs())
}

func main() {
	var a Abser // nil

	// describe(a)

	var p *Vertex // nil
	// p = &Vertex{1, 2}

	a = p

	describe(a)
}
