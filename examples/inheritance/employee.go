package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return "Name: " + p.Name + " | Age: " + strconv.Itoa(p.Age)
}

type Employee struct {
	Person
	Designation      string
	YearOfExperience int
}

func (e Employee) String() string {
	// return "Name: " + e.Name + " | Age: " + strconv.Itoa(e.Age) + " | Designation: " + e.Designation + " | Experience: " + strconv.Itoa(e.YearOfExperience)
	return e.Person.String() + " | Designation: " + e.Designation + " | Experience: " + strconv.Itoa(e.YearOfExperience)
}

func main() {
	p := Person{"Gaurav", 29}

	emp := Employee{p, "GA", "Principal Consultant", 8}

	fmt.Println(emp.Person)

	emp.Person.Name = "G A"
	emp.Name = "algogrit"

	fmt.Println(emp.Person)
	fmt.Println(p)

	// fmt.Println(emp.Person.String())
	// fmt.Println(emp.String())

	// fmt.Println(emp)
}
