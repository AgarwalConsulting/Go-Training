package main

import "fmt"

func describeType(s Shape) {
	fmt.Printf("%#v, %T\n", s, s)
}

func main() {
	var s Shape

	describeType(s)

	s = Square{10}

	describeType(s)

	fmt.Println("Area of square:", s.Area())

	c := Cube{10}

	s = c

	describeType(s)

	c1, _ := s.(Cube) // Type inference & casting

	fmt.Println("Surface area of cube:", s.Area())
	fmt.Println("Volume of cube:", c1.Area())
}
