package main

import (
	"fmt"
	"unsafe"
)

type Node struct {
	Element int
	Float   float32
	Flag    bool
}

func main() {
	a := [4]int{}
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(Node{}))
}
