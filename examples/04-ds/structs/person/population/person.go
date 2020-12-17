package population

import (
	"time"

	"github.com/AgarwalConsulting/Go-Training/live-examples/location"
)

type Person struct {
	Name        string
	Age         int
	yearOfBirth int
	Address     location.Address
}

func NewPerson(name string, age int, city string) Person {
	yearOfBirth := time.Now().Year() - age

	return Person{name, age, yearOfBirth, location.Address{City: city}}
}

func WithName(person Person, name string) Person {
	person.Name = name
	return person
}
