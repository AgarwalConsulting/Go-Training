package main

import (
	"fmt"

	"github.com/AgarwalConsulting/Go-Training/examples/packages/simple-package/mypkg"
)

func main() {
	fmt.Println("Hello,", mypkg.Greeting())
	// fmt.Println(mypkg.value)
	fmt.Println(mypkg.Value)
}
