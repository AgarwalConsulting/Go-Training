package main

import "fmt"

type Person struct {
	FirstName, LastName, Name string
}

func main() {
	var p *Person

	fmt.Println(p) // nil

	p = new(Person)

	p.FirstName = "Gaurav"

	fmt.Printf("%#v\n", p)
}
