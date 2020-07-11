package main

import "fmt"

type Person struct {
	age  *int   // nil
	name string // ""
}

func newPerson() Person {
	person := Person{name: "GA"}

	return person
}

func main() {
	person := newPerson()

	fmt.Println(person.name)
	fmt.Println(person.age)

	person.age = new(int)

	fmt.Println(person.age)

	page := person.age

	*page = 28

	fmt.Println(person)
	fmt.Println(*person.age)
}
