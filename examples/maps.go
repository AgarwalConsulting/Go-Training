package main

import "fmt"

type Place string

type Vertex struct {
	X, Y float64
}

var m map[Place]Vertex = map[Place]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func (place Place) findThePlace() Vertex {
	return m[place]
}

func main() {
	google := Place("Synamedia")
	fmt.Println(google.findThePlace())
}
