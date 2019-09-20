package Shapes

import "math"

type Circle struct {
	Radius float64
}

// This method has a receiver type of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}
