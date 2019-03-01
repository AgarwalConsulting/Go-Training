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
	quit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()

	go fibonacci(c)

	<-quit
}
