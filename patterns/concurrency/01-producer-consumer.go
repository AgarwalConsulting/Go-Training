package main

import "fmt"

func fibonacci(ch chan<- int) {
	x, y := 0, 1
	for {
		ch <- x
		x, y = y, x+y
	}
}

func main() {
	c := make(chan int)
	// quit := make(chan bool)

	// n := 10

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	go fibonacci(c)

	for {

	}
	// <-quit
}
