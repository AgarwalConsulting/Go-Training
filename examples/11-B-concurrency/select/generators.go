package main

import (
	"fmt"
	"time"
)

func sendPeriodically(period time.Duration, ch chan<- int) {
	counter := 0
	for {
		time.Sleep(period)
		counter++
		ch <- counter // Sends value to the channel once every given duration
	}
}

func generator(n time.Duration) chan int {
	ch := make(chan int)

	fmt.Println("Produces value every: ", n)

	go sendPeriodically(n, ch)

	return ch
}

func main() {
	ch1 := generator(time.Microsecond * 970123) // 970.12ms
	ch2 := generator(time.Microsecond * 80000)  // 80ms

	var x int

	// for {
	// 	x = <-ch1 // Block until generator 1 sends a value
	// 	fmt.Println("Received for generator 1: ", x)
	// 	x = <-ch2 // Block until generator 2 sends a value
	// 	fmt.Println("Received for generator 2: ", x)
	// }

	// for {
	// 	select {
	// 	case x = <-ch1:
	// 		fmt.Println("Received from 1st generator")
	// 	case x = <-ch2:
	// 		fmt.Println("Received from 2nd generator")
	// 	}
	// 	fmt.Println("Received: ", x)
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
