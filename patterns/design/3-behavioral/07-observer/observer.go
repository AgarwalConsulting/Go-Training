
// Observer => Notified of the change

type Subscriber interface {
	Notify(change interface{})
}

// Observable => Produces
