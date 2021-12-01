package main

import "fmt"

type Engine struct{}

func (Engine) Start() {
	fmt.Println("Starting...")
}

func (Engine) Stop() {
	fmt.Println("Stopping...")
}

// Car has-a Engine
type Car struct {
	Engine
}
