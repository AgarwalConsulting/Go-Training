package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int) // Unbuffered channel

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Sending...")
		ch <- 42 // Sending to a channel is going to block until there is a receive
		fmt.Println("Sent!")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Receiving...")
		time.Sleep(time.Second * 3)
		val := <-ch // Receiving from a channel is going to block until there is a send
		fmt.Println("Recieved: ", val)
	}()

	wg.Wait()

	// fmt.Println("Next: ", <-ch)
}
