package main

import "math"

type Circle float64

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(float64(*c), 2)
}

func (c *Circle) Scale() {
	*c = *c * 10
}
