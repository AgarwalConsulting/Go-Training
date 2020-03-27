package main

import "fmt"

type Place struct {
	Name     string
	Location string
}

type Vertex struct {
	X, Y float64
}

var m map[Place]Vertex = map[Place]Vertex{
	{"Bell Labs", "E Coast"}: Vertex{
		40.68433, -74.39967,
	},
	{"Google", "Mountain View"}: Vertex{
		37.42202, -122.08408,
	},
}

func (place Place) findThePlace() Vertex {
	return m[place]
}

// place name => place lat, lng => query google maps reverse geolocation API => Address

// func getLocation(place) chan<- Vertex
// func getAddress(<-chan Vertex) chan<- Address

func main() {
	fmt.Println("Implement a streaming and processing pipeline")
}
