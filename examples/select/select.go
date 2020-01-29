package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	timer := time.After(time.Millisecond * 1)
	for {
		select {
		case c <- x: // send
			x, y = y, x+y
		case <-timer: // receive
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	fibonacci(c, quit)
}
