package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
}

func main() {
	p := Person{"G A", 29}

	fmt.Println(p)
}
