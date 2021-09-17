package main

import (
	"fmt"
	"time"
)

func another() {
	panic("Oops...")
}

func main() {
	go another()

	time.Sleep(1 * time.Second)

	fmt.Println("Does this print?")
}
