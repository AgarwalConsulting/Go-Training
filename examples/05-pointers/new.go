package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var p *Vertex

	// p = new(Vertex)
	fmt.Printf("%p\n", p)

	if p != nil {
		fmt.Println("X: ", p.X)
	}
}
