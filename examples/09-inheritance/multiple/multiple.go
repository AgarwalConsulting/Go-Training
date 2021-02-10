package main

import "fmt"

type Animal struct {
	legs int
}

func (a Animal) Move() {
	fmt.Printf("Animal moves on %d legs\n", a.legs)
}

type Bird struct {
	wings int
}

func (b Bird) Move() {
	fmt.Printf("Bird flies using %d wings\n", b.wings)
}

type Lion struct {
	Animal
}

type Ostrich struct {
	Animal
	Bird
}

func main() {
	l := Lion{Animal{4}}

	l.Move()

	o := Ostrich{Animal{2}, Bird{2}}

	o.Move()
}
