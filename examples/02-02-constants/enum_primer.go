package main

import (
	"fmt"
)

func printStatus(s status) {
	fmt.Println("The status is:", s)
}

type status string

const (
	running status = "running"
	waiting status = "blocked"
)

func main() {
	var s status

	s = running

	fmt.Println(s) // Default value

	fmt.Printf("%v, %T\n", s, s)
}
