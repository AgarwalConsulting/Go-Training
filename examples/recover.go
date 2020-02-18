package main

import (
	"fmt"
)

func panicky() {
	defer func() {
		fmt.Println("Defered...")

		recover()
		// if r := recover(); r != nil {
		// 	fmt.Printf("Recovered: %#v, %T\n", r, r)
		// }
	}()

	defer panic("oh no!")

	i := 0

	fmt.Println(100 / i)

	fmt.Println("Did I panic?")
}

func main() {
	defer fmt.Println("Defered to main...")

	panicky()

	fmt.Println("Hello, World!")
}
