package main

import "fmt"

type Hello struct {
	x *int
	y string
}

func hello() *Hello {
	// var hello Hello
	// hello.x = 42
	// hello.y = "hello"

	// num := 42

	hi := Hello{y: "hello"}

	return &hi
}

func main() {
	// var hello Hello = hello()

	mainHi := hello()
	var x **int
	x = &mainHi.x
	xDeref := mainHi.x

	fmt.Printf("%T, %#v\n", mainHi.x, mainHi.x)

	// *x = 100
	hundred := 100
	mainHi.x = &hundred

	fmt.Printf("%T, %#v\n", mainHi, mainHi)
	fmt.Printf("%T, %#v\n", mainHi.x, mainHi.x)

	fmt.Printf("%T, %#v\n", x, x)
	fmt.Printf("%T, %#v\n", xDeref, xDeref)
	// fmt.Printf("%T, %#v\n", y, y)

	fmt.Printf("%f, %T\n", float64(**x), float64(**x))
}
