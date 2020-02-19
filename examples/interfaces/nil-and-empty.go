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
	if v == nil {
		return 0
	}

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser

	fmt.Printf("%T, %v\n", a, a) // nil interface

	// a.Abs() // panic: runtime error: invalid memory address or nil pointer dereference

	if a == nil {
		fmt.Println("nil!")
	}

	var v *Vertex // nil

	a = v

	fmt.Printf("%T, %v\n", a, a) // empty interface

	a.Abs()

	if a == nil {
		fmt.Println("nil!")
	}
}
