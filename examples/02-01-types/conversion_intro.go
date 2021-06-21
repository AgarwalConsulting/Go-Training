package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// var i int
	// var s string
	// var f float64

	var (
		i int
		f float64
		s string
	)

	fmt.Printf("%#v, %#v, %#v\n", i, f, s)

	i = 42

	f = float64(i)

	// s = string(i)
	s = strconv.Itoa(i)

	fmt.Printf("%T, %#v\n", i, i)
	fmt.Printf("%T, %#v\n", f, f)
	fmt.Printf("%T, %#v\n", s, s)

	fmt.Println("----")

	var k int

	f = 42.12
	k = int(f)

	fmt.Printf("Float: %T, %#v\n", k, k)

	// k = int(s)
	// var err error
	k, err := strconv.Atoi(s)

	fmt.Println("Error: ", err)
	fmt.Printf("String: %T, %#v\n", k, k)

	s = "1234as"
	l, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("unable to convert: ", err)
		os.Exit(1)
	}

	fmt.Printf("String: %T, %#v\n", l, l)
}
