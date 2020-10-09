package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

// func (p Person) String() string {
// 	return "Name: " + p.Name + " | Age: " + strconv.Itoa(p.Age)
// }

func main() {
	p := Person{"G A", 29}

	fmt.Println(p)
}
