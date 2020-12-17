package main

import (
	"fmt"

	"github.com/AgarwalConsulting/Go-Training/live-examples/location"
	"github.com/AgarwalConsulting/Go-Training/live-examples/population"
)

func main() {
	var gaurav population.Person

	fmt.Println(gaurav) // {"" 0 0 {"" "" ""}}
	fmt.Printf("%#v\n", gaurav)

	gaurav = population.Person{
		Name:    "G A",
		Age:     29,
		Address: location.Address{City: "Chennai"},
	} // Struct literal syntax

	// gaurav = population.NewPerson("G A", 29, "Chennai")

	// fmt.Println(gaurav.yearOfBirth)

	fmt.Printf("%T\n", gaurav)
	fmt.Println(gaurav)

	fmt.Println(gaurav.Name)
	fmt.Println(gaurav.Address.City)
}
