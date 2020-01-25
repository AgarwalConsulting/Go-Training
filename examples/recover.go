package main

import (
	"fmt"
)

func panicky() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %#v, %T\n", r, r)
		}
	}()

	i := 0

	fmt.Println(100 / i)
	// panic("oh no!")
}

func main() {
	panicky()

	fmt.Println("Hello, World!")
}
