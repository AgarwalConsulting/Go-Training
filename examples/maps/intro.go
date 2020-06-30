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

	// m := map[string]Vertex{}

	m["Google"] = Vertex{37.42202, -122.08408}
	m["Alphabet"] = Vertex{37.42202, -122.08408}

	// m := map[string]Vertex{
	// 	"Google":    {37.42202, -122.08408},
	// 	"Bell Labs": {40.68433, -74.39967},
	// }

	fmt.Println(m)

	fmt.Println(m["Google"])

	// fmt.Println(m["AcmeInc"]) // How do you know it's default?

}
