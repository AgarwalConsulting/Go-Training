package main

import (
	"fmt"
	"strconv"
	"time"
)

type MessageMetadata struct {
	Timestamp   string `json:"timestamp"`
	MessageType string `json:"messageType"`
}

type InfoPayload struct {
	MessageMetadata
	Info string `json:"info"`
}

func printMessageType(m any) {
	v := m.(MessageMetadata) // ?

	fmt.Println(v.MessageType)
}

func main() {
	p := InfoPayload{
		MessageMetadata: MessageMetadata{
			Timestamp:   strconv.Itoa(int(time.Now().Unix())),
			MessageType: "UserEvent",
		},
		Info: "Lorem Ipsum",
	}

	printMessageType(p)
}
