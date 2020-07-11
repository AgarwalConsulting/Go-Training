// https://tour.golang.org/flowcontrol/12
package main

import (
	"fmt"
)

func deferedPanic() {
	panic("oh no!")
}

func say(hi string) {
	defer fmt.Println("world")
	defer fmt.Println("3")
	defer fmt.Println("2")
	// defer deferedPanic()
	defer fmt.Println("1")

	// var i, j int

	// fmt.Println(i / j)

	// os.Exit(1)

	fmt.Println(hi)
}

func main() {
	hello := "hello"
	defer fmt.Println("Exit main!")

	fmt.Println("Init main...")

	say(hello)

	fmt.Println("Say returned!")
}
