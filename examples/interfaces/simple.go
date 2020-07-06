package main

import "fmt"

type MyFloat float64

func (f MyFloat) String() string {
	real := int(f)
	fraction := f - MyFloat(real)

	return fmt.Sprintf("Real: %d, Fraction: %f", real, fraction)
}

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old.", p.name, p.age)
}

func main() {
	f := MyFloat(64.2)

	fmt.Println(f)

	p := Person{"Goper", 10}

	fmt.Printf("%v\n", p) //
	// PrintVerbose(f)
}
