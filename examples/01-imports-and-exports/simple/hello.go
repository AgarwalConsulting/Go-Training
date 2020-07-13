package main

import (
	"fmt"

	"github.com/AgarwalConsulting/Go-Training/examples/01-imports-and-exports/simple/extras"
	"github.com/AgarwalConsulting/Go-Training/examples/01-imports-and-exports/simple/mypkg"
)

func main() {
	fmt.Println(extras.Hello, ",", mypkg.Greeting())
	// fmt.Println(mypkg.value)
	fmt.Println(mypkg.Value)
}
