package pubsub

import "sync"

// PubSub: Many Consumers & One or Many Producer
// .Subscribe() to register a listener
// .Send to send a value to all the listener

type Data interface{}

// type Subscriber func(Data)

type Subscriber interface {
	Notify(Data)
}

// SubscriberFn is a concrete type which implements Subscriber interface
// it can wrap a func(Data)
type SubscriberFn func(Data)

func (fn SubscriberFn) Notify(value Data) {
	fn(value)
}

// PubSub uses the Observer pattern to notify multiple subscribers
type PubSub struct {
	Subscribers []Subscriber
	Mut         sync.RWMutex
	Channel     chan Data
}

// Subscribe registers a Subscriber to get notified on new data
func (ps *PubSub) Subscribe(sub Subscriber) {
	ps.Mut.Lock()
	ps.Subscribers = append(ps.Subscribers, sub)
	ps.Mut.Unlock()
}

// Send has a fire & forget semantics to notify all the subscribers
func (ps *PubSub) Send(value Data) {
	ps.Channel <- value
}

func New() *PubSub {
	ps := PubSub{}
	ps.Channel = make(chan Data)

	go func() {
		for {
			value := <-ps.Channel
			ps.Mut.RLock()
			for _, subscriber := range ps.Subscribers {
				// subscriber.Notify(value)
				go subscriber.Notify(value)
			}
			ps.Mut.RUnlock()
		}
	}()

	return &ps
}
