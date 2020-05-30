package main

import (
	"fmt"
	"time"
)

func fibGenerator(ch chan<- int, dur time.Duration) {
	i, j := 0, 1

	tch := time.After(dur)

	for {
		fmt.Println("Sending...", i)
		select {
		case ch <- i:
			time.Sleep(100 * time.Millisecond)
			i, j = j, i+j
		case <-tch:
			fmt.Println("Time up!")
			fmt.Println("Exiting...")
			goto exit
		}
	}
exit:
	close(ch)
}

func main() {
	ch := make(chan int)

	go fibGenerator(ch, 1*time.Second)

	for x := range ch {
		fmt.Println("Received", x)
		time.Sleep(100 * time.Millisecond)
	}
}
