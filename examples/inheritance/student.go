package main

import (
	"fmt"
	"strconv"
)

type Detailer interface {
	String() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return "Name: " + p.Name + " Age: " + strconv.Itoa(p.Age)
}

type Student struct {
	Person // Person Person
	Marks  []float64
}

func PrintDetails(d Detailer) {
	fmt.Println(d)
}

func main() {
	p := Person{"Gaurav", 28}
	student := Student{p, []float64{100, 50, 20, 80}}

	PrintDetails(p)
	fmt.Println(student) // ?
}
