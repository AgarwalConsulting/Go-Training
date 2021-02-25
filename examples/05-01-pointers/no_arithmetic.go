package main

import "fmt"

func main() {
	var a = [4]int{2, 3, 5, 7}

	var p *int
	// p = a
	p = &a

	p1 := p + 1 // no pointer arithematic

	fmt.Println(*p1)
}
