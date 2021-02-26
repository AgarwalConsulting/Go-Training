package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processIncomingRequest(id int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from:", r)
		}
	}()

	if rand.Int()%2 == 0 {
		panic("some error")
	}

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
