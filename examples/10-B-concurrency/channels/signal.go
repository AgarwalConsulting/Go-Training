package main

import (
	"fmt"
	"time"
)

func fibonacciGen(q chan bool) <-chan int {
	i, j := 0, 1
	ch := make(chan int, 10)

	// Producer
	go func() {
		for {
			i, j = j, j+i
			nextVal := j - i

			select {
			case <-q:
				fmt.Println("Quitting...")
				close(ch)
				return
			case ch <- nextVal:
				fmt.Println("Sending...", nextVal)
			}
		}
	}()

	return ch
}

func main() {
	quit := make(chan bool)
	fibChan := fibonacciGen(quit)
	counter := 0

	for el := range fibChan {
		time.Sleep(time.Second * 1)

		fmt.Println("Next Value: ", el)
		counter++

		if counter == 10 {
			quit <- true
			// close(fibChan)
			// Stop the fibonacciGen
		}
	}
}
