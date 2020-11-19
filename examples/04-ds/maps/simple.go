package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	var m map[string]Vertex

	if m == nil {
		fmt.Println("Maps are nil by default!")
	}

	m = make(map[string]Vertex)
	// m := map[string]Vertex{} // Map literal syntax

	m["Google"] = Vertex{37.42202, -122.08408}
	m["Alphabet"] = Vertex{37.42202, -122.08408}
	m["Bell Labs"] = Vertex{37.123, -52.08408}

	// m = map[string]Vertex{
	// 	"Google":    {37.42202, -122.08408},
	// 	"Alphabet":  {37.42202, -122.08408},
	// 	"Bell Labs": {40.68433, -74.39967},
	// }

	for key, value := range m {
		fmt.Println("Coordinates for", key, ":", value)
	}

	v := m["Google"]
	fmt.Println("Google is at:", v)

	x := m["AcmeInc"]
	fmt.Println("Acme is at: ", x)

	// delete(m, "Google")

	// fmt.Println(m)
}
