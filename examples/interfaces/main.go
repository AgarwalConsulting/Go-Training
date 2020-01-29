package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Abser interface {
	Abs() float64
}

func returnRandomAbser() Abser {
	f := MyFloat(-math.Sqrt2) // MyFloat

	var v *Vertex     // nil
	v = &Vertex{3, 4} // *Vertex

	if rand.Int()%2 != 0 {
		return f
	}

	return v
}

func main() {
	var a Abser

	a = returnRandomAbser()

	if f, ok := a.(MyFloat); ok {
		fmt.Println("It is a MyFloat!", f)
	}

	if _, ok := a.(*Vertex); ok {
		fmt.Println("It is a vertex pointer!")
	}

	fmt.Println(a.Abs())
}
