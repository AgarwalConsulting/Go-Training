package main

import "fmt"

func createSequence() func() int {
	var seq int
	seq = 1000

	return func() int {
		seq += 1

		return seq
	}
}

func main() {
	f := createSequence()

	for i := 0; i < 10; i++ {
		fmt.Println(f()) //
	}

	f1 := createSequence()

	fmt.Println(f1()) //
}
