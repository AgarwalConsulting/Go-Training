package main

import "math"

type Cube float64

func (c Cube) Area() float64 {
	return float64(6 * c * c)
}

func (c Cube) Volume() float64 {
	return math.Pow(float64(c), 3)
}
