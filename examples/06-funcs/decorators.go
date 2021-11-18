package main

import (
	"fmt"
	"log"
)

func logDecorator(op func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		log.Println("Input:", a, b)
		output := op(a, b)
		log.Println("Output:", output)
		return output
	}
}

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Result:", Add(10, 12))

	AddWithLogging := logDecorator(Add)

	fmt.Println("Result:", AddWithLogging(10, 12))
}
