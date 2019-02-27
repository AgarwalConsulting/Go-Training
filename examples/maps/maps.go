package main

import (
	"fmt"
	"strings"
)

type Place struct {
	Name     string
	Location string
}

type Vertex struct {
	X, Y float64
}

var m map[string]Vertex = map[string]Vertex{
	Place{"Bell Labs", "E Coast"}.hashCode(): Vertex{
		40.68433, -74.39967,
	},
	Place{"Google", "Mountain View"}.hashCode(): Vertex{
		37.42202, -122.08408,
	},
}

func (place Place) hashCode() string {
	locationStr := strings.Split(place.Location, " ")

	var locationInitial string

	for _, val := range locationStr {
		locationInitial = locationInitial + string([]rune(val)[0])
	}

	return fmt.Sprintf("%s/%s", place.Name, locationInitial)
}

func (place Place) findThePlace() Vertex {
	return m[place.hashCode()]
}

func main() {
	google := Place{"Google", "M V"}

	// fmt.Println(google.hashCode())
	fmt.Println(google.findThePlace())
}
