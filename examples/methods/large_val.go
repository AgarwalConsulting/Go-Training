package main

import "fmt"

type Names [100000]string

func (names *Names) Print() {
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	var n Names

	fmt.Println(n)

	n.Print()
}
