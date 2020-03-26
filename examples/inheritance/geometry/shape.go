package main

type Shape interface {
	Area() float64
	Scale(float64)
}

type Shape3D interface {
	Shape
	Volume() float64
}
