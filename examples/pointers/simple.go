package main

import "fmt"

func greeting(s *string) {
	*s = "Hello, " + *s
}

func main() {
	s := "World!"

	var sp *string

	fmt.Println(sp)

	greeting(&s)

	fmt.Println(s)
}
