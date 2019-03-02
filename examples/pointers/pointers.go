package main

import (
	"fmt"
)

func say(hi *string) {
	*hi += ", world!"
}

func main() {
	hello := "hello"

	say(&hello)

	fmt.Println(hello)
}
