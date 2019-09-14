package main

import (
	"fmt"
	"unsafe"
)

type Empty struct{}

func main() {
	e := Empty{}
	fmt.Println(unsafe.Sizeof(e))
}
