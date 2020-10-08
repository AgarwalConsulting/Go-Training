package main

import "fmt"

type Person struct {
	FirstName, LastName string
	Name                *string
}

func combineName(s Person) {
	*s.Name = s.FirstName + " " + s.LastName
}

func main() {
	person := Person{FirstName: "Gaurav", LastName: "Agarwal", Name: new(string)}

	fmt.Println(*person.Name) // ""

	combineName(person)

	fmt.Println(*person.Name) // "G A"
	fmt.Printf("%#v\n", person)
}
