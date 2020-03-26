package main

import "fmt"

type Any interface{}

func print(a Any) {
	// v, ok := a.(string)

	// if ok {
	// 	fmt.Println("a is a string of val: ", v)
	// }

	switch a.(type) {
	case int:
		fmt.Println(a, " is int")
	case string:
		fmt.Println(a, " is string")
	default:
		fmt.Println(a, " is unknown type!")
	}
}

func main() {
	var a Any

	a = 42
	describe(a)
	print(a)
	a = false
	describe(a)
	print(a)
	a = "Hello"
	describe(a)
	print(a)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
