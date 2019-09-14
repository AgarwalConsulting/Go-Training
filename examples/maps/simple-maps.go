package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	m := map[string]Vertex{
		"Google":    {37.42202, -122.08408},
		"Bell Labs": {40.68433, -74.39967},
	}

	fmt.Println(m)

	fmt.Println(m["Google"])

	v, ok := m["GlobalEdge"]

	if ok {
		fmt.Println(v)
	}

	m["GlobalEdge"] = Vertex{12, 43.21}

	m["Google"] = Vertex{}

	for k, v := range m {
		fmt.Printf("Key: %s; Value: %#v\n", k, v)
	}

	fmt.Println("Deleting...")

	delete(m, "Google")

	fmt.Printf("Length: %d\n", len(m))

	for k, v := range m {
		fmt.Printf("Key: %s; Value: %#v\n", k, v)
	}
}
