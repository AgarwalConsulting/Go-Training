package main

import "fmt"

func main() {
	var a Abser

	fmt.Println(a) // nil

	// a = MyFloat(42)

	v := Vertex{3, 5}

	v.Scale(2)

	fmt.Println(v)

	a = &v

	fmt.Println(a, a.Abs())
}
