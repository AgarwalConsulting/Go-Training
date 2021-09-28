package main

import "fmt"

func main() {
	var p1 struct {
		a int
		b string
	}

	var p2 struct {
		a int
		b string
	}

	p1 = p2
	p2 = p1

	fmt.Println(p1, p2)
	fmt.Println("Is equal?", p1 == p2)
}
