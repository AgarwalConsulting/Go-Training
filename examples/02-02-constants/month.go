package main

import (
	"fmt"
	"time"
)

func main() {
	currentMonth := time.Now().Month()

	prevMonth := currentMonth - 1

	nextMonth := currentMonth + 1

	fmt.Println("Prev month was: ", prevMonth)
	fmt.Println("Looking forward to: ", nextMonth)
}
