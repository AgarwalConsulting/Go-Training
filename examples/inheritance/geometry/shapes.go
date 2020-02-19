package main

type Shape interface {
	Area() float64
}

type ScalableShape interface {
	Shape
	Scale()
}

type Shape3D interface {
	Shape
	Volume() float64
}
