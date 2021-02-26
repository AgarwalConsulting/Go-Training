package main

import (
	"fmt"
	"math/rand"
	"time"
)

func performTask(id int, retry bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)

			if retry == true {
				fmt.Println("\tRetrying...", id)
				performTask(id, false)
			}
		}
	}()

	if rand.Int()%2 == 0 {
		panic("some error")
	}
}

func processIncomingRequest(id int) {
	// defer recover()
	performTask(id, true)

	fmt.Println("Processed", id, "successfully!")
}

func main() {
	rand.Seed(time.Now().Unix())
	i := 0
	for {
		go processIncomingRequest(i)
		time.Sleep(1 * time.Second)
		i++
	}
}
