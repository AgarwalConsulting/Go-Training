package main

import "fmt"

type MyInt int

func main() {
	var v MyInt

	i := 42 // Type: int

	// v = 42
	// v = i // Will this work?

	v = MyInt(i)

	// var f float32

	// f := float32(42.10) // Conversion

	// fmt.Printf("%T\n", f) // float32

	// v := MyInt(42) // Conversion

	fmt.Printf("%T %v\n", v, v) // int
}
