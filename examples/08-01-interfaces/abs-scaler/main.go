package main

import (
	"fmt"

	"algogrit.com/mtdex/myfloat"
	"algogrit.com/mtdex/vertex"
)

type Abser interface {
	Abs() float64
}

type AbsoluteScaler interface {
	Abs() float64
	Scale(float64)
}

func printAbsolute(a Abser) {
	fmt.Printf("%T %#v\n", a, a)

	fmt.Println("\tAbsolute value:", a.Abs())
}

func printAbsoluteAndScale(a AbsoluteScaler) {
	fmt.Printf("%T %#v\n", a, a)
	fmt.Println("\tAbsolute value:", a.Abs())

	a.Scale(2)

	fmt.Println("\tScaled Absolute value:", a.Abs())
}

func main() {
	f := myfloat.MyFloat(-42)

	printAbsolute(f)
	// printAbsoluteAndScale(f)

	pf := &f // Type: *myfloat.MyFloat

	printAbsolute(pf)
	// printAbsoluteAndScale(pf)

	v := vertex.Vertex{X: 3, Y: 4}

	printAbsolute(v) // ?
	// v.Scale(2)
	// printAbsoluteAndScale(v)

	pv := &v

	printAbsolute(pv)
	printAbsoluteAndScale(pv)
}
