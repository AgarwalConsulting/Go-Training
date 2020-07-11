package main

import "fmt"

func main() {
	var p *string

	fmt.Println(p) // nil

	fmt.Println(*p) // ?
}
