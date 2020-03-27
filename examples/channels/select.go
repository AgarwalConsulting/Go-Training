package main

import (
	"fmt"
	"time"
)

func sequence(n int, sleepDuration time.Duration, ch chan<- int) {
	init := n
	noOfValuesToSend := 1000

	for {
		if n < init + noOfValuesToSend {
			time.Sleep(sleepDuration)
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
	ch1 := make(chan int, 5)
	ch2 := make(chan int)

	go sequence(100000, time.Millisecond * 100, ch1)
	go sequence(1000, time.Millisecond * 50, ch2)

	for {
		select {
		case seq1Val := <-ch1:
			fmt.Println("Received from first sequence: ", seq1Val)
		case seq2Val := <-ch2:
			fmt.Println("Received from second sequence: ", seq2Val)
		}
		time.Sleep(time.Millisecond * 10)
	}
}
