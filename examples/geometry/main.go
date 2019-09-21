package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/geometry/shapes"
)

func describe(s *shapes.Shape) {
	fmt.Printf("%T %v\n", s, s)
}

func doublesAndPrintsArea(shapeName string, s *shapes.Shape) {
	describe(s)
	if s != nil {
		// s.Double()
		fmt.Println("Area of ", shapeName, ": ", (*s).Area())
	}
}

func main() {
	var c shapes.Shape
	c = &shapes.Circle{10}

	var ptr *shapes.Shape
	ptr = &c

	// var sq shapes.Shape
	// var rect *shapes.Shape = &shapes.Rectangle{}

	var sqPtr *shapes.Square // nil
	var sq shapes.Shape = sqPtr
	// sq = &shapes.Square{10}

	doublesAndPrintsArea("circle", ptr)
	// doublesAndPrintsArea("rectangle", rect)
	doublesAndPrintsArea("square", sq)
}
