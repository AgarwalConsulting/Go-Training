package main

import (
	"fmt"
)

type Place struct {
	Name string
}

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m map[Place]Vertex

	if m == nil {
		fmt.Println("Maps are nil by default!")
	}

	fmt.Printf("%T\n", m) // map[main.Place]main.Vertex

	m = make(map[Place]Vertex)

	// s := "Google"

	m[Place{"Google"}] = Vertex{37.42202, -122.08408}

	fmt.Printf("%#v\n", m)

	v, ok := m[Place{"Google"}]

	if ok {
		fmt.Println("Google is at: ", v)
	} else {
		fmt.Println("Unable to locate Google!")
	}
}
