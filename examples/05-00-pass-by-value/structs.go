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

func updateAddress(p *Place, newAddress string) {
	p.Address = newAddress
}

func relocateLat(p Place, newLat float64) {
	p.Location.Lat = newLat
}

func main() {
	var acme Place
	fmt.Println(acme) // {"" nil ""}

	acme = Place{
		"Acme Inc.",
		&Vertex{34.42, 53.5612},
		"Mars",
	}

	updateAddress(&acme, "Venus")
	fmt.Println(acme.Address)

	relocateLat(acme, 42.12)
	fmt.Println(acme.Location.Lat) //
}
