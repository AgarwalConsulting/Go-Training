package main

import (
	"fmt"
)

type status int

// const (
// 	running status = 0
// 	waiting status = 1
// 	done    status = 2
// )

const (
	running status = iota
	waiting
	done
)

func main() {
	var s status

	fmt.Println(s) // Default value: 0

	fmt.Printf("%v, %T\n", running, running)

	fmt.Printf("%v, %T\n", done, done)
}
