package main

import "math"

// Vertex represents a point in plane. *Vertex implements Abser interface.
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
