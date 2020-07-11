package main

import "fmt"

func greeting(s string) {
	s = "Hello, " + s
}

func main() {
	s := "World!"

	greeting(s)

	fmt.Println(s)
}
