package main

import (
	"fmt"
)

func print(p *int) {
	fmt.Println(p)  // address of a
	fmt.Println(*p) // value of a
}

func main() {
	a := 42 // int

	print(&a) // *int
}
