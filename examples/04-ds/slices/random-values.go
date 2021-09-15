package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getNumber() int {
	return rand.Intn(100)
}

func main() {
	// var values [100]int
	var values []int // Slice

	fmt.Printf("%T %v\n", values, values)
	fmt.Println("\tLen:", len(values), "Cap:", cap(values), "Nil?", values == nil)
	fmt.Printf("\tAddress: %p\n", values)

	// values = []int{} // Empty slice
	// values = make([]int, 100) //
	values = make([]int, 0, 100) //

	fmt.Printf("%T %v\n", values, values)
	fmt.Println("\tLen:", len(values), "Cap:", cap(values), "Nil?", values == nil)
	fmt.Printf("\tAddress: %p\n", values)

	i := 0
	for {
		num := getNumber()
		fmt.Println("Values:", len(values), "Next Number:", num)
		fmt.Println("\tLen:", len(values), "Cap:", cap(values), "Nil?", values == nil)
		fmt.Printf("\tAddress: %p\n", values)

		time.Sleep(100 * time.Millisecond)

		values = append(values, num)
		// values[i] = num

		if num == 42 {
			break
		}
		i++
	}

	fmt.Println(values)
}
