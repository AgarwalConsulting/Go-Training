package main

import (
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest"})

	if err != nil {
		log.Fatal(err)
	}

	var run bool = true
	for run == true {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			// application-specific processing
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
}
