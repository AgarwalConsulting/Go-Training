package main

import "fmt"

type MyInt int

func main() {
	var m MyInt

	fmt.Printf("%T %v\n", m, m) // main.MyInt, 0

	m = 42 // Works!

	i := 42

	m = i // Will not work!

	// m = MyInt(i)

	fmt.Println(m)
}
