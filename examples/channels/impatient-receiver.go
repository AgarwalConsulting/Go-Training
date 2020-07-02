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
	var c chan int // nil
	c = make(chan int, 10)

	go Producer(c, 100)

	flag := true

	for flag {
		select {
		case nextVal := <-c:
			fmt.Println("Processed: ", nextVal)
		case <-time.After(10 * time.Second):
			fmt.Println("Tired of waiting")
			flag = false
		}
	}
}
