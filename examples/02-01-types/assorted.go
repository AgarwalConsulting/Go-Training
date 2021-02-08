package main

import "fmt"

func main() {
	// i, j, k, l, m, n := "hello", 32, 10.12, 42.0, true, 10i

	var (
		i string
		j int
		k float64
		l float64
		m bool
		n complex128
	)

	fmt.Printf("%T %#v\n", i, i) // string, ""
	fmt.Printf("%T %#v\n", j, j) // int, 0
	fmt.Printf("%T %#v\n", k, k) // float64, 0.0
	fmt.Printf("%T %#v\n", l, l) // float64, 0.0
	fmt.Printf("%T %#v\n", m, m) // bool, false
	fmt.Printf("%T %#v\n", n, n) // complex128, 0+0i

	fmt.Println("----")

	i, j, k, l, m, n = "hello", 32, 10.12, 42.0, true, 10i

	fmt.Printf("%T %#v\n", i, i) // string
	fmt.Printf("%T %#v\n", j, j) // int
	fmt.Printf("%T %#v\n", k, k) // float64
	fmt.Printf("%T %#v\n", l, l) // float64
	fmt.Printf("%T %#v\n", m, m) // bool
	fmt.Printf("%T %#v\n", n, n) // complex128

	// j = 200000000000

	// fmt.Printf("%T %v\n", j, j)

	// var-of-type = type(var-of-other-type)
	j = int(l)

	fmt.Printf("%T %v\n", j, j)
}
