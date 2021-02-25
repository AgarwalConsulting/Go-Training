package main

import "fmt"

func A() {
	defer fmt.Println("Hello from A!") //
	defer B()

	fmt.Println("Exiting A...") //
}

func B() {
	defer fmt.Println("Hello from B!") //
	defer A()

	fmt.Println("Exiting B...") //
}

func main() {
	defer B()
}
