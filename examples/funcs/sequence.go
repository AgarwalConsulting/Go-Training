package main

import "fmt"

func createSequence() func() int {
	seq := 1000

	return func() int {
		seq += 1

		return seq
	}
}

func main() {
	f := createSequence()
	f1 := createSequence()

	for i := 0; i < 10; i++ {
		fmt.Println(f()) // 1001 - 1010
	}

	fmt.Println(f1()) // 1001
}
