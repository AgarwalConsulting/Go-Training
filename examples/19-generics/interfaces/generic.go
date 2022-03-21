package main

import (
	"fmt"
	"math"
)

type Shape[T any] interface {
	Area() T
}

type Square struct {
	side int
}

func (s Square) Area() int {
	return s.side * s.side
}

type Circle struct {
	radius int
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

// type Number interface{ int | float64 }

func printArea(s Shape[int]) {
	fmt.Printf("%T | %#v\n", s, s)
	fmt.Println("\tArea:", s.Area())
}

func main() {
	sq := Square{2}

	printArea(sq)

	var sh Shape[float64]

	c := Circle{3}

	sh = c

	fmt.Println(sh)
	// printArea(c)
}
