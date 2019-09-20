package Shapes

import "fmt"

type Shape interface {
	Area() float64
	Circumference() float64
}

func PrintArea(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Printf("Shape (%T: %v) \tArea: %v \tCircumference: %v\n", s, s, s.Area(), s.Circumference())
	}
}
