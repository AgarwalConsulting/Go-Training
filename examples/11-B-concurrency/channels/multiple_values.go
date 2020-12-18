package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) { // Send-only channel
	i, j := 0, 1

	for {
		fmt.Println("Sending...")
		ch <- i
		i, j = j, i+j
		fmt.Println("Sent!")
	}
}

func main() {
	var ch chan int

	go producer(ch)

	fmt.Println("Receiving...")
	for i := 0; i < 10; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println("Received: ", <-ch)
	}
}
