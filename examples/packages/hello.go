package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/packages/geo"
)

func main() {
	p := person{"G", 28}

	fmt.Println(geo.Location{12.31, 42.12})

	fmt.Println("Hello, World!", p)
}
