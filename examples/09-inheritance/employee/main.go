package main

import "fmt"

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
