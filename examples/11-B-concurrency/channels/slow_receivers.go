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
			// close(ch)
			fmt.Println("Quitting...")
			return
		}
		n += 1
	}
}

func main() {
	ch := make(chan int) // Unbuffered
	go sequence(10, ch)

	for x := 0; x < 10; x++ {
		el := <-ch
		fmt.Println("Received:", el)
		time.Sleep(3 * time.Second)
	}
}
