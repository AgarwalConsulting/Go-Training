package shapes

import "math"

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Double() {
	c.Radius = c.Radius * 2
}
