package main

import (
	"fmt"
	"time"
)

func Producer(c chan<- int, n int) {
	i, j := 0, 1
	for x := 0; x < n; x++ {
		i, j = j, i+j
		nextVal := j - i
		c <- nextVal // Sending blocks only if the buffer is full
		fmt.Println("Sent: ", nextVal)

		dur := time.Second * time.Duration(nextVal)
		fmt.Println("Sleeping for... ", dur)
		time.Sleep(dur)
	}
	close(c)
}

func main() {
	var c chan int         // nil
	c = make(chan int, 10) // Buffered

	go Producer(c, 100)

	// for x := range c {
	// 	// If while receiving any value takes more than 5 seconds, then terminate the program.
	// 	fmt.Println("Received: ", x)
	// }

	for {
		x, ok := <-c
		if !ok {
			return
		}
		fmt.Println("Received: ", x)
	}
}
