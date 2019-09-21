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
