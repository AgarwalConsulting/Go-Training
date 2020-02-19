package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

// Abs is a receiver function
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser

	fmt.Printf("%T, %v\n", a, a) // nil interface

	var v *Vertex // nil

	a = v

	fmt.Printf("%T, %v\n", a, a) // empty interface

	if a == nil {
		fmt.Println("nil!")
	}
}
