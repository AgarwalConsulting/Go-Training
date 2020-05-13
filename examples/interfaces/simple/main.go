package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Scale(float64)
}

type Outline interface {
	Perimeter() float64
}

func PrintPerimeter(o Outline) {
	fmt.Println("Perimeter is: ", o.Perimeter())
}

func PrintAreaAndScale(shapeName string, s Shape) {
	if s == nil {
		fmt.Println(shapeName, "is nil")
		return
	}
	fmt.Println("Area of a", shapeName, "is: ", s.Area())
	s.Scale(2)
	fmt.Println("Scaled area is: ", s.Area())
}

func main() {
	var s Shape

	fmt.Println("----Nil Interface-----")

	fmt.Printf("%T, %v\n", s, s)

	PrintAreaAndScale("Shape", s)

	var p *Circle // nil

	fmt.Println("----Interface with value as nil----")

	s = p

	PrintAreaAndScale("Circle", p)

	fmt.Printf("%T, %v\n", s, s)

	// c := Circle{12.0}
	// p := &c // *Circle

	// s = p

	// PrintAreaAndScale("Circle", s)

	// fmt.Println("Area of a circle is: ", p.Area())

	// r := Rectangle{2, 5}

	// s = &r

	// PrintPerimeter(r)
	// PrintAreaAndScale("Rectangle", s)

	// sq := Square{5}

	// PrintAreaAndScale("Square", &sq)
}
