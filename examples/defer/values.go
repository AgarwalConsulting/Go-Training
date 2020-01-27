package main

import "fmt"

func main() {
	var i = 42

	defer fmt.Println("Deferred value of i:", i)

	i = 10
}
