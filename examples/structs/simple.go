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

type Person struct {
	Name string
	Age  int
	Address
}

func main() {
	var p Person
	// p := Person{}
	// p := Person{"Gaurav", 29}
	// p := Person{Name: "Gaurav"}

	fmt.Println(p.Address) // {"" {0 0}}
}
