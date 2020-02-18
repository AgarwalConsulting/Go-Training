package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// var m map[string]Vertex

	// m = make(map[string]Vertex)

	// m["Google"] = Vertex{37.42202, -122.08408}
	// m["Alphabet"] = Vertex{37.42202, -122.08408}

	m := map[string]Vertex{
		"Google":    {37.42202, -122.08408},
		"Bell Labs": {40.68433, -74.39967},
	}

	fmt.Println(m)

	fmt.Println(m["Google"])

	// fmt.Println(m["AcmeInc"]) // How do you know it's default?

	v, ok := m["AcmeInc"]

	if ok {
		fmt.Println(v)
	}

	m["AcmeInc"] = Vertex{12, 43.21}

	fmt.Println(m)

	m["Google"] = Vertex{}

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("Key: %s; Value: %#v\n", k, v)
	}

	fmt.Println("Deleting...")

	delete(m, "Google")

	fmt.Printf("Length: %d\n", len(m))
	// fmt.Printf("Capacity: %d\n", cap(m)) // -- Doesn't work!
}
