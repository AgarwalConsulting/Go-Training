package main

import (
	"fmt"
	"time"
)

func fibGenerator(c chan<- int) {
	i, j := 0, 1
	for x := 0; x < 8; x++ {
		i, j = j, i+j // 1, 0+1
		nextVal := j - i
		fmt.Printf("Sleeping for %ds...\n", nextVal)
		time.Sleep(time.Second * time.Duration(nextVal))
		c <- nextVal // Sending blocks only if the buffer is full
		fmt.Println("Sent: ", nextVal)
	}
	fmt.Println("Completed generator!")
	close(c)
}

func Consumer(c <-chan int) {
	for {
		fmt.Println("Processing...")
		i, ok := <-c
		if ok {
			fmt.Println("Processed: ", i)
		} else {
			return
		}
	}
}

func main() {
	var c chan int // nil
	c = make(chan int, 4)

	go fibGenerator(c)

	Consumer(c)
}
