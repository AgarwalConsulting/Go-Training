package main

import (
	"fmt"

	"algogrit.com/geometry/shapes"
)

func printShapeAndArea(name string, sh shapes.Shape) {
	fmt.Printf("%T %v\n", sh, sh)
	fmt.Println("\tArea of", name, "is: ", sh.Area())

	// <interface-var>.(Type)
	rect, ok := sh.(*shapes.Rectangle) // Type Assertion

	if ok {
		fmt.Println("\tPerimeter is ", rect.Perimeter())
	}
}

func printShape3DAndArea(name string, sh shapes.Shape3D) {
	fmt.Printf("%T %v %p\n", sh, sh, sh)
	fmt.Println("\tArea of", name, "is: ", sh.Area())
	fmt.Println("\tVolume of", name, "is: ", sh.Volume())
}

func main() {
	s := shapes.Square{5.0}
	printShapeAndArea("square", &s) // string, *Square

	r := shapes.Rectangle{2, 4}
	printShapeAndArea("rectangle", &r) // string, Rectangle
	// fmt.Println(r)

	c := shapes.NewCircle(3.0)
	printShapeAndArea("circle", &c)

	cu := shapes.Cube{7.0}
	printShape3DAndArea("cube", &cu)
	// printShapeAndArea("cube", &cu)
}
