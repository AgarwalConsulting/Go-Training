package main

import (
	"fmt"
)

func addGreeting(hi *string) {
	fmt.Printf("%d\n", hi)
	*hi += ", world!"
}

func main() {
	hello := "hello"

	addGreeting(&hello)

	fmt.Println(hello)
}
