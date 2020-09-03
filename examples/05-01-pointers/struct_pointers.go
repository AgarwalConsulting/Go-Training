package main

import "fmt"

type Person struct {
	FirstName, LastName, Name string
}

func combineName(s *Person) {
	s.Name = s.FirstName + " " + s.LastName
}

func main() {
	person := Person{FirstName: "Gaurav", LastName: "Agarwal"}

	fmt.Println(person.Name) // ""

	combineName(&person)

	fmt.Println(person.Name)    // ""
	fmt.Printf("%#v\n", person) // ""
}
