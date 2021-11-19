package main

import "fmt"

func fib() func() int {
	i, j := -1, 1
	return func() int {
		i, j = j, i+j
		if j > 42 {
			panic("some error")
		}
		return j
	}
}

func start() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exiting gracefully...", r)
		}
	}()
	f := fib()
	for {
		fmt.Println(f())
	}
}

func main() {
	for {
		start()
	}
}
