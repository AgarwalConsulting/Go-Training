package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Legs() int { return 4 }

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs\n", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int { return 5 }

func main() {
	var octo OctoCat
	fmt.Println(octo.Legs()) // 5
	octo.PrintLegs()         // I have 4 legs
}

// Instead describe Specification pattern
// Implement a product with size and color attributes
