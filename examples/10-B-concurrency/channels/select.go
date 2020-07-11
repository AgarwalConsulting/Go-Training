package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator(n time.Duration) chan int {
	counter := 0
	ch := make(chan int)

	fmt.Println("Sleep for ", n)

	go func() {
		for {
			time.Sleep(n)
			counter++
			ch <- counter // Sends value to the channel once every given duration
		}
	}()

	return ch
}

func main() {
	ch1 := generator(time.Duration((rand.Int() % 1000000000)))
	ch2 := generator(time.Duration((rand.Int() % 1000000000)))

	var x int

	// for {
	// 	x = <-ch1 // Block until generator 1 sends a value
	// 	fmt.Println("Received for generator 1: ", x)
	// 	x = <-ch2 // Block until generator 2 sends a value
	// 	fmt.Println("Received for generator 2: ", x)
	// }
	for {
		newValue := false
		select {
		case x = <-ch1:
			fmt.Println("Received from 1st generator")
			newValue = true
		case x = <-ch2:
			fmt.Println("Received from 2nd generator")
			newValue = true
		default:
			fmt.Println(".")
			time.Sleep(10 * time.Millisecond)
			newValue = false
		}
		if newValue {
			fmt.Println("Received: ", x)
		}
	}
}
