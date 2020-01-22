package main

import "fmt"

func main() {
	v := "true" // change me!

	var any interface{}

	any = v

	if val, ok := any.(bool); ok {
		fmt.Printf("%T\n", val)
		fmt.Println("Yay! It's a bool!")
	}

	fmt.Printf("v is of type %T\n", v)
}
