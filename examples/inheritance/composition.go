package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Name() string {
	return p.name
}

type Student struct {
	Person
	totalMarks int
}

func main() {
	s := Student{Person{"G", 28}, 100}

	fmt.Println(s.name)

	fmt.Println(s.Name())
}
