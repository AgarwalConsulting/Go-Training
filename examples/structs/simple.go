package main

import "fmt"

type Vertex struct {
	Lat  float64
	Long float64
}

type Address struct {
	Location    string
	Coordinates Vertex
}

func createPerson() *Person {
	return new(Person)
}

type Person struct {
	Name string
	Age  int
	Address
	// Address Address
}

func main() {
	p := Person{"Gaurav", 29, Address{"Chennai", Vertex{42.12, 12.12}}}

	// fmt.Println("Person is: ", p)

	fmt.Println("Coordinates: ", p.Address.Coordinates)
}
