package shapes

import "math"

type Circle struct {
	radius float64
}

func NewCircle(radius float64) Circle {
	return Circle{radius: radius}
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Scale(by float64) {
	c.radius = c.radius * by
}
