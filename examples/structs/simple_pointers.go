package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Person struct {
	Name     string
	Age      int
	Location Vertex
}

func updatePerson(p *Person, newName string) {
	p.Name = newName
}

func main() {
	person := Person{"G A", 28, Vertex{42, 12}}

	updatePerson(&person, "Gaurav")

	fmt.Printf("%#v\n", person)
}
