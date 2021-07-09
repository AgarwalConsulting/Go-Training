package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Deep Copy: Design a Person with Address
// Copy Method: Implement a DeepCopy on `Person` & `Address` types
// Copy through serialization: Use `bytes` & `encoding/gob`
// Prototype Factory: Employee with Name & Office fields; Design NewFunctions with different locations

type Address struct {
	Street string
	City   string
	State  string
}

// func (a *Address) Clone() Address {
// 	a1 := *a // a1 is a copy of a

// 	// a1 := Address{}
// 	// a1.Street = a.Street
// 	// a1.City = a.City
// 	// a1.State = a.State

// 	return a1
// }

type Person struct {
	Name    string
	Address *Address
}

func (p *Person) Clone() Person {
	b := bytes.Buffer{}

	gob.NewEncoder(&b).Encode(p)

	var p1 Person

	gob.NewDecoder(&b).Decode(&p1)

	return p1
}

func main() {
	p := Person{Name: "John", Address: &Address{Street: "1st Street", City: "New York", State: "NY"}}

	// p1 := p // p1 is a copy of p
	p1 := p.Clone() // p1 is a copy of p

	p1.Name = "Jane"
	p1.Address.Street = "2nd Street"
	p1.Address.City = "San Francisco"
	p1.Address.State = "CA"

	fmt.Println(p, p1)

	fmt.Printf("%#v %#v", p, p1)
	fmt.Printf("%#v %#v", p.Address, p1.Address)
}
