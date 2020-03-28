package main

import (
	"fmt"
	"sync"
	"time"

	"algogrit.com/pubsub/pkg/pubsub"
)

type Tracker struct {
	Values []pubsub.Data
	Count  int
	Mut    sync.RWMutex
}

func (av *Tracker) Notify(value pubsub.Data) {
	av.Mut.Lock()
	av.Values = append(av.Values, value)
	av.Count++
	av.Mut.Unlock()
}

func (av *Tracker) String() string {
	av.Mut.RLock()
	defer av.Mut.RUnlock()

	return fmt.Sprintf("Values: %s, Count: %d", av.Values, av.Count)
}

func main() {
	ps := pubsub.New()

	subFn1 := pubsub.SubscriberFn(func(value pubsub.Data) {
		fmt.Println("Received new data in Subscriber 1: ", value)
	})

	subFn2 := pubsub.SubscriberFn(func(value pubsub.Data) {
		fmt.Println("Received new data in Subscriber 2: ", value)
	})

	tracker := &Tracker{} // Type: *Tracker

	ps.Subscribe(subFn1)
	ps.Subscribe(tracker)

	ps.Send(pubsub.Data("Hello, World!")) // This needs to be async

	ps.Subscribe(subFn2)

	ps.Send(pubsub.Data(false))

	fmt.Println("Successfully send data to PubSub!")

	fmt.Println(tracker) // &{["Hello, World!"] 1}

	time.Sleep(1 * time.Millisecond)

	fmt.Println(tracker)

	for {

	}
}
