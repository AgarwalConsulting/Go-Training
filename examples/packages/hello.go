package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/packages/geo"
)

func main() {
	p := person{"G", 28}

	loc := geo.Location{12.31, 42.12}

	// loc.Print()

	geo.Print(loc)

	fmt.Println("Hello, World!", p)
}
