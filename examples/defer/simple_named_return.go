package main

import "fmt"

func c() (i int) {
	defer func() { i++ }()
	i = 10
	return 1
}

func main() {
	fmt.Println(c())
}
