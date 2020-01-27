package main

import (
	"fmt"
	"time"
)

type status int

const (
	single status = iota
	engaged
	married
)

func main() {
	// s := status(1)
	// day := time.Weekday(7) % 7

	// fmt.Println("The status is: ", s)
	// fmt.Printf("%v, %T\n", s, s)
	// fmt.Println("The day is: ", day)
	// fmt.Printf("%v, %T\n", day, day)

	fmt.Println(time.Duration(1000000))
}
