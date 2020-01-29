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
			ch <- counter
		}
	}()

	return ch
}

func main() {
	ch1 := generator(time.Duration((rand.Int() % 1000000000)))
	ch2 := generator(time.Duration((rand.Int() % 1000000000)))

	var x int
	for {
		select {
		case x = <-ch1:
			fmt.Println("Recived from 1st generator")
		case x = <-ch2:
			fmt.Println("Recived from 2nd generator")
		}
	}

	fmt.Println("Received: ", x)
}
