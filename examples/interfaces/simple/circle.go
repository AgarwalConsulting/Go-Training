package main

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	if c != nil {
		return math.Pi * c.Radius * c.Radius
	}
	return 0
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}

func (c *Circle) Scale(scaleBy float64) {
	if c != nil {
		c.Radius = c.Radius * scaleBy
	}
}
