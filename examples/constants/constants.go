package main

import (
	"fmt"
	"time"
)

type status int

const (
	running status = iota
	waiting
	done
)

func main() {
	var s status

	fmt.Println(s) // Default value

	fmt.Printf("%v, %T\n", running, running)

	fmt.Println(time.Microsecond * 10000)

	fmt.Println(time.Weekday(0))
}
