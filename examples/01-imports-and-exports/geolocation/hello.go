package main

import (
	"fmt"

	geo "github.com/AgarwalConsulting/Go-Training/examples/01-imports-and-exports/geolocation/geo"
	"github.com/AgarwalConsulting/Go-Training/examples/01-imports-and-exports/geolocation/geo/reverse"
)

func main() {
	p := person{"G", 28}

	loc := geo.Location{12.31, 42.12}

	// loc.Print()

	geo.Print(loc)

	fmt.Println("Hello, World!", p)

	fmt.Println(reverse.Reverse())
}
