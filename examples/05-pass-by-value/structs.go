package main

import "fmt"

type Place struct {
	Name     string
	Location *Vertex
	Address  string
}

type Vertex struct {
	Lat, Long float64
}

func (p Place) Relocate(v Vertex) {
	p.Location = &v
}

func (p Place) RelocateLat(v float64) {
	p.Location.Lat = v
}

func (p Place) String() string {
	return fmt.Sprintf("Name: %s, Location: %#v, Address: %s", p.Name, p.Location, p.Address)
}

func main() {
	ge := Place{
		"GE",
		&Vertex{34.42, 53.5612},
		"Whitefield, Bengaluru",
	}

	fmt.Println(ge)

	ge.Relocate(Vertex{0, 0})

	fmt.Println(ge)

	ge.RelocateLat(0)

	fmt.Println(ge)
}
