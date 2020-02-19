package main

import "fmt"

func PrintShapeAndArea(shapeName string, c ScalableShape) {
	c.Scale()
	fmt.Println("Area of a ", shapeName, ": ", c.Area())
}

func PrintShapeAreaAndVolume(shapeName string, c Shape3D) {
	fmt.Println("Surface area of a ", shapeName, ": ", c.Area(), "& volume: ", c.Volume())
}

func main() {
	s := Square(10)
	// r := Rectangle{10, 20}
	c := Circle(5)

	PrintShapeAndArea("Square", &s)
	// PrintShapeAndArea("Rectangle", r)
	fmt.Println(c.Area())
	PrintShapeAndArea("Circle", &c)

	PrintShapeAreaAndVolume("Cube", Cube(9))
}
