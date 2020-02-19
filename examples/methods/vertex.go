package main

import (
	"math"
)

// Vertex is a user-defined type
type Vertex struct {
	X, Y float64
}

// Abs is a receiver function
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(s float64) {
	v.X = v.X * s
	v.Y = v.Y * s
}
