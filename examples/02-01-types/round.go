package main

import (
	"fmt"
	"math"
)

func main() {
	a := 42.5
	b := 43.5

	fmt.Println(math.Round(a), math.Round(b))
	fmt.Println(math.RoundToEven(a), math.RoundToEven(b))
}
