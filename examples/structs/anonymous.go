package main

import "fmt"

func main() {
	var p struct {
		Name string
		Age  int
	}

	p = struct {
		Name string
		Age  int
	}{"GA", 29}

	fmt.Printf("%T\n", p)
	fmt.Println(p)
}
