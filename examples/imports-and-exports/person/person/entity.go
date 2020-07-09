package person

import "time"

var Hello = "hello"

type Person struct {
	name        string
	age         int
	yearOfBirth int
}

func New(name string, age int) Person {
	yearOfBirth := time.Now().Year() - age
	return Person{name, age, yearOfBirth}
}
