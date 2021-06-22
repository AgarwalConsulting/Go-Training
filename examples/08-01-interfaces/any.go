package main

import "fmt"

type Any interface {
	SomeFn() string
}

type Person struct {
	Name string
}

func (p Person) SomeFn() string {
	return p.Name
}

func main() {
	var a Any

	p := Person{"Gaurav"}

	a = p

	describe(a)
}

func describe(i Any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
