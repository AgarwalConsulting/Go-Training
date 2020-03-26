package main

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Scale(scaleBy float64) {
	c.Radius = c.Radius * scaleBy
}
