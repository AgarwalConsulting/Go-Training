package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/interfaces/simple/square"
)

type Shape interface {
	Area() float64
	Scale(float64)
}

type Outline interface {
	Perimeter() float64
}

func printShapeAndArea(shapeName string, s Shape) {
	fmt.Println(shapeName, s)
	fmt.Println("Area: ", s.Area())
	s.Scale(2)
	fmt.Println("Scaled Area: ", s.Area())
}

func main() {
	var o Outline

	c := Circle{5} // Value type of Circle
	printShapeAndArea("Circle", &c)

	sq := square.Square{7} // Square
	printShapeAndArea("Square", &sq)

	r := Rectangle{10, 5}
	printShapeAndArea("Rectangle", &r)

	o = r
	fmt.Println("Perimeter is:", o.Perimeter())

	// p := &r // *Rectangle

	// p.Area()
	// p.Scale(2)
	// p.Perimeter()

	// r.Scale(2) // ?
	// r.Area()
	// r.Perimeter()

	// p.Scale(2)

	// // r.Scale(2)

	// fmt.Println(r.Area())
}
