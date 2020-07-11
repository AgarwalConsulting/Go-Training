package main

import (
	"fmt"
	"time"
)

func sequence(noOfValuesToSend int, ch chan<- int) {
	init := 1000
	n := 0

	for {
		if n < noOfValuesToSend {
			nextVal := init + n
			fmt.Println("Sending: ", nextVal)
			ch <- nextVal // Blocks if there is no receiver (unbuffered) or channel is full (buffered)
		} else if noOfValuesToSend == n {
			close(ch)
		}
		n += 1
	}
}

func main() {
	ch := make(chan int, 10)
	go sequence(10, ch)

	for el := range ch { // Loops for next value received from channel; Blocks until the channel is closed
		fmt.Println("Received:", el)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Range is completed!")

	nextVal, ok := <-ch // Receiving from a closed channel doesn't block and gives the zero value of type

	if !ok {
		fmt.Println("Channel is closed: ", nextVal)
	}
}
