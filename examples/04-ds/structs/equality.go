package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{"Gaurav", 30}

	p1 := struct {
		age  int
		name string
	}{30, "Gaurav"}

	fmt.Println(p == p1)

	p = p1
	p1 = p

	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", p1)
}
