package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/packages/simple-package/mypkg"
)

func main() {
	fmt.Println("Hello,", mypkg.Greeting())
	// fmt.Println(mypkg.value)
	fmt.Println(mypkg.Value)
}
