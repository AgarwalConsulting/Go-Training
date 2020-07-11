package main

import "fmt"

func update(s string) {
	s += " world"
}

func main() {
	s := "hello"
	update(s)

	fmt.Println(s)
}
