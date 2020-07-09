package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func updateGoogle(m map[string]Vertex) {
	m["Google"] = Vertex{}
}

func main() {
	m := map[string]Vertex{
		"Google":    {37.42202, -122.08408},
		"Bell Labs": {40.68433, -74.39967},
	}

	updateGoogle(m)

	fmt.Println(m)
}
