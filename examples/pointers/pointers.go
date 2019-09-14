package main

import (
	"fmt"
)

func addGreeting(hi *string) {
	*hi += ", world!"
}

func main() {
	hello := "hello"

	addGreeting(&hello)

	fmt.Println(hello)
}
