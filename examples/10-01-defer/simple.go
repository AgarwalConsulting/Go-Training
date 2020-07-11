package main

import "fmt"

func say() string {
	fmt.Println("Saying...")
	return "world!"
}

func main() {
	defer fmt.Println(say())

	fmt.Println("hello")
}
