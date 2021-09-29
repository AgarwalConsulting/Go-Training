package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func getData(id int) <-chan int {
	c := make(chan int)

	go func() {
		fmt.Println("doing some processing...", id)

		valuesToSend := rand.Intn(100)

		for i := 0; i < valuesToSend; i++ {
			fmt.Println("\t\tValues to send:", valuesToSend, "from:", id, " | Values sent:", i)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			fmt.Println("\tSending data from:", id, "...")

			// data := rand.Int()
			c <- id // Sending only once
		}
		close(c)
	}()

	return c
}

func merge(chs ...<-chan int) <-chan int {
	fmt.Printf("%T\n", chs) // []<-chan int

	outputCh := make(chan int)

	var wg sync.WaitGroup
	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for data := range ch {
				outputCh <- data
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(outputCh)
	}()

	return outputCh
}

func main() {
	rand.Seed(time.Now().Unix())

	a := getData(1)
	b := getData(2)

	c := merge(a, b)

	for x := range c {
		fmt.Println("Received from:", x)
	}
}
