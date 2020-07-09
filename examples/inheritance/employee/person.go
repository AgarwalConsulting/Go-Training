package main

import "strconv"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return "Name: " + p.Name + " | Age: " + strconv.Itoa(p.Age)
}
