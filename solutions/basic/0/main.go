package main

import (
	"fmt"

	"github.com/AgarwalConsulting/Go-Training/solutions/basic/0/fizzbuzz"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fizzbuzz.FizzBuzz(i + 1))
	}
}
