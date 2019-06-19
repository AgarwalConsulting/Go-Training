package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m["Synamedia"] = Vertex{12, 43.21}

	keys := []string{"Cisco", "Tata", "Google"}

	for _, key := range keys {
		v, ok := m[key]

		fmt.Printf("%#v\n", ok)

		if ok {
			fmt.Println("key: ", key, " value: ", v)
		}
	}

	// for key, value := range m {
	// 	fmt.Println("key: ", key, " value: ", value)
	// }

	fmt.Println(m)
}
