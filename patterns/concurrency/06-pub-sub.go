package main

import (
    "fmt"
	"sync"
	"time"
)

type Subscriber func(int)

type PubSub struct {
    c chan int
    Subscribers []Subscriber
    lock sync.RWMutex
}

func (pubSub *PubSub) Subscribe(subscriber Subscriber) {
    pubSub.lock.Lock()
    pubSub.Subscribers = append(pubSub.Subscribers, subscriber)
    pubSub.lock.Unlock()
}

func (pubSub *PubSub) Send(data int) {
    pubSub.c <- data
}

func CreateNewPubSub() *PubSub {
    var pubSub PubSub
    pubSub.c = make(chan int)

    go func(pubSub *PubSub) {
        for {
            data := <- pubSub.c

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
    pubSub := CreateNewPubSub()

    pubSub.Subscribe(func(data int) {
        fmt.Println("Received new data: ", data)
    })

    pubSub.Send(42)

    time.Sleep(1 * time.Millisecond)
}
