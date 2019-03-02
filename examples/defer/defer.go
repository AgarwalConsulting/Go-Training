// https://tour.golang.org/flowcontrol/12
package main

import (
	"fmt"
)

func say(hi string) {
	defer fmt.Println("world")
	defer fmt.Println("3")
	defer fmt.Println("2")
	defer fmt.Println("1")

	// panic("oh no!")
	// os.Exit(0)

	fmt.Println(hi)
}

func main() {
	hello := "hello"
	defer fmt.Println("Exit main!")

	fmt.Println("Init main...")

	say(hello)
}
