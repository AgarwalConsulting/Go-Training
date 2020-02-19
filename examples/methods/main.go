package main

import "fmt"

type Abser interface {
	Abs() float64
}

type Scalar interface {
	Scale(float64)
}

func ScaleAndPrint(val Scalar) {
	val.Scale(10)
	fmt.Printf("%T, %v\n", val, val)
}

func PrintValueAndAbs(val Abser) {
	fmt.Println(val, val.Abs())
}

func main() {
	v := Vertex{3, 4}
	f := MyFloat(-12.12)

	ScaleAndPrint(&v)

	PrintValueAndAbs(v)
	PrintValueAndAbs(f)
}
