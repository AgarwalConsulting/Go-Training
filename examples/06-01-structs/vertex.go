package main

import "fmt"

// type Vertex struct {
// 	X int
// 	Y int
// }

// func createVertex() Vertex {
// 	return Vertex{10, 20}
// }

func myPrint(any struct {
	X int
	Y int
}) {
	fmt.Printf("%T, %#v, %v\n", any, any, any)
}

func main() {
	type Vertex struct {
		X int
		Y int
	}

	v := Vertex{10, 20}

	myPrint(v)

	// v := createVertex()

	p := &v

	fmt.Println(v)
	fmt.Println(p.X)
	fmt.Println(p.Y)

	fmt.Printf("%T, %#v, %v\n", p, p, p)
}
