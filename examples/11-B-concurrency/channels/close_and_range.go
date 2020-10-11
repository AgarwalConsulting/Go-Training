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
			ch <- nextVal // Blocks if there is no receiver (unbuffered) or channel is full (buffered)
			fmt.Println("Sent: ", nextVal)
		} else if noOfValuesToSend == n {
			close(ch) // Closing a channel is not neccessary!
			fmt.Println("Quitting...")
			return
		}
		n += 1
	}
}

func main() {
	ch := make(chan int, 5) // Buffered
	go sequence(10, ch)

	// for x := 0; x < 10; x++ {
	// 	el := <-ch
	// 	fmt.Println("Received:", el)
	// 	time.Sleep(3 * time.Second)
	// }

	for el := range ch { // Loops for next value received from channel; Blocks until the channel is closed
		fmt.Println("Received:", el)
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Range is completed!")

	// nextVal := <-ch // Type: int
	// fmt.Println("Next Value:", nextVal)

	nextVal, ok := <-ch // Receiving from a closed channel doesn't block and gives the zero value of type

	if !ok {
		fmt.Println("Channel is closed: ", nextVal)
	}

	// ch <- 42 // Sending to a closed channel blocks
}
