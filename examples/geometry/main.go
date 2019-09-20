package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/geometry/shapes"
)

func doublesAndPrintsArea(shapeName string, s shapes.Shape) {
	fmt.Printf("%#v\n", s)
	s.Double()
	fmt.Println("Area of ", shapeName, ": ", s.Area())
}

func main() {
	c := shapes.Circle{10}
	sq := shapes.Square{10}

	doublesAndPrintsArea("circle", &c)
	doublesAndPrintsArea("square", &sq)
}
