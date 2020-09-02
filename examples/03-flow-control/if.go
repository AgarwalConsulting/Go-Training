package main

import (
	"fmt"
	"math"
)

func main() {
	// i := math.Sqrt(25)
	if i := math.Sqrt(25); i < 6 {
		fmt.Println("i is: ", i)
	} else {
		fmt.Println("i is large!")
	}

	// fmt.Println(i)
}
