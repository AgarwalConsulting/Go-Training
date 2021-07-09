package main

import "fmt"

// Multiple Aggregation: Bird & Lizard types with Age fields combine into a Dragon; Introduce Aged interface
// Decorator: Shape interface with Render method; Circle & Square types; Make them have color as well; Introduce ColoredShape and OpacityShape type

func greetingPrinter(printFn func(string)) func(string) {
	return func(name string) {
		fmt.Print("Hello, ")
		printFn(name)
	}
}

func print(name string) {
	fmt.Println("My name is:", name)
}

func main() {
	greetingPrinter(print)("Gaurav")
}
