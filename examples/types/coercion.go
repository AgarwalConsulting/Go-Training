package main

import (
	"fmt"
	"strconv"
)

func main() {
	v := "true" // change me!

	var any interface{}

	any = v

	// Inference
	if val, ok := any.(bool); ok {
		fmt.Printf("%T\n", val)
		fmt.Println("Yay! It's a bool!")
	}

	fmt.Printf("v is of type %T\n", v)

	// Coercion
	num := 42.0 // float

	numint := int(num)

	fmt.Printf("%#v, %T\n", numint, numint)

	// Conversion
	str := "64"

	if i, err := strconv.Atoi(str); err != nil {
		fmt.Println("Conversion: ", i)
	}

	// fmt.Println("Conversion: ", i)
}
