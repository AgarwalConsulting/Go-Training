// Use strings.Builder
// Builder: Design a html builder
// Builder Facet: Design a PersonBuilder, PersonJobBuilder, PersonAddressBuilder
// Builder Parameter: Design an EmailBuilder => func SendEmail(action func(b *EmailBuilder) {})
// Functional Builder: Design PersonBuilder combining Facet with Builder Parameter for lazy building

package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Job     Job
	Address Address
}

type PersonBuilder struct {
	person Person
}

func (pb *PersonBuilder) Build(fn func(*PersonBuilder)) {
	fn(pb)
}

func (pb *PersonBuilder) Name(name string) *PersonBuilder {
	pb.person.Name = name
	return pb
}

func (pb *PersonBuilder) WithJob() *JobBuilder {
	return &JobBuilder{PersonBuilder: pb}
}

func (pb *PersonBuilder) WithAddress() *AddressBuilder {
	return &AddressBuilder{PersonBuilder: pb}
}

type Job struct {
	Designation string
	StartDate   string
	Company     string
}

type JobBuilder struct {
	*PersonBuilder
	// job Job
}

func (ab *JobBuilder) Designation(designation string) *JobBuilder {
	ab.person.Job.Designation = designation
	return ab
}

type Address struct {
	DoorNumber string
	Street     string
	City       string
	Pincode    string
}

type AddressBuilder struct {
	*PersonBuilder
	// address Address
}

func (ab *AddressBuilder) City(name string) *AddressBuilder {
	ab.person.Address.City = name
	return ab
}

func main() {
	pb := PersonBuilder{}

	pb.Build(func(pb *PersonBuilder) {
		pb.Name("Gaurav").
			WithJob().
			Designation("Developer").
			WithAddress().
			City("Chennai")
	})

	fmt.Println(pb.person)
}
