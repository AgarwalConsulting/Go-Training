package main

import "fmt"

// Producer
func fibonacci(ch chan<- int, quit <-chan struct{}) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:

			x, y = y, x+y
		case <-quit:
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan struct{})

	// Consumer
	go func() {
		for i := 0; i < 20; i++ {
			fmt.Println(<-c)
		}
		quit <- struct{}{}
	}()

	go fibonacci(c, quit)

	for {

	}
	// <-quit
}
