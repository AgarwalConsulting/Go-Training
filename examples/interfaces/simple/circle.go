package main

import "math"

// Circle implements Shape
type Circle struct {
	Radius float64
}

// Area is a value receiver on Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Scale is a pointer receiver on Circle
func (c *Circle) Scale(scaleBy float64) {
	c.Radius = c.Radius * scaleBy
}
