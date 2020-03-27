package main

import "fmt"

func DescribeAndGetArea(s Shape) {
	fmt.Printf("%T, %#v\n", s, s)

	fmt.Printf("Area of %T: %v is %f\n", s, s, s.Area())
	s.Scale(2)
	fmt.Printf("Scaled Area of %T: %v is %f\n", s, s, s.Area())
}

func main() {
	var shape Shape // Type: Shape which is a interface

	fmt.Println(shape.Area())
	// DescribeAndGetArea(shape)

	var square *Square // Value: nil

	fmt.Printf("%T, %v\n", square, square)

	// fmt.Println(square.Area())

	DescribeAndGetArea(square)

	// s := Square{10}

	// DescribeAndGetArea(&s) // 100

	// c := Circle{7}

	// DescribeAndGetArea(&c)

	// r := Rectangle{10, 5}

	// DescribeAndGetArea(&r)
}
