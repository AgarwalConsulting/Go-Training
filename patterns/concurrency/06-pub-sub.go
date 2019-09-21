package main

import (
	"fmt"
	"sync"
	"time"
)

type Subscriber func(int)

type PubSub struct {
	c           chan int
	Subscribers []Subscriber
	lock        sync.RWMutex
}

func (pubSub *PubSub) Subscribe(subscriber Subscriber) {
	pubSub.lock.Lock()
	pubSub.Subscribers = append(pubSub.Subscribers, subscriber)
	pubSub.lock.Unlock()
}

func (pubSub *PubSub) Publish(data int) {
	pubSub.c <- data
}

func CreatePubSub() *PubSub {
	var pubSub PubSub
	pubSub.c = make(chan int)

	go func(pubSub *PubSub) {
		for {
			data := <-pubSub.c

			pubSub.lock.RLock()
			for _, subscriber := range pubSub.Subscribers {
				subscriber(data)
			}
			pubSub.lock.RUnlock()
		}
	}(&pubSub)

	return &pubSub
}

func main() {
	pubSub := CreatePubSub()

	pubSub.Subscribe(func(data int) {
		fmt.Println("Received new data: ", data)
	})

	pubSub.Publish(42)

	go pubSub.Subscribe(func(data int) {
		fmt.Println("Subscriber 2: ", data)
	})

	go pubSub.Publish(10)

	fmt.Println("Publishing asynchronously?")

	time.Sleep(1 * time.Second)
}
