package main

import (
	"fmt"
	"time"
)

func sequence(n int, ch chan<- int) {
	init := n
	noOfValuesToSend := 1

	for {
		if n < init + noOfValuesToSend {
			fmt.Println("Sending: ", n)
			ch <- n // Blocks if there is no receiver (unbuffered) or channel is full (buffered)
		} else if n == init + noOfValuesToSend {
			close(ch)
			// return
			// ch <- n
		}
		n += 1
	}
}

func main() {
	ch := make(chan int, 5)

	go sequence(100000, ch)

	openVal, ok := <-ch // Blocks if there is nothing to received: no sender (unbuffered) or channel is empty (buffered)

	fmt.Println("Channel is open: ", openVal, ok)

	for el := range ch { // Loops for next value received from channel; Blocks until the channel is closed
		fmt.Println("Received:", el)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Range is completed!")

	closedVal, ok := <-ch

	fmt.Println("Channel is closed: ", closedVal, ok)
}
