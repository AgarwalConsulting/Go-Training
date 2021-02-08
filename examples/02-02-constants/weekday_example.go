package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	day := now.Weekday()

	// day := time.Weekday(5)

	fmt.Printf("%T %#v\n", day, day) //
	fmt.Println(day)                 //

	tomorrow := day + 1
	fmt.Printf("%T %#v\n", tomorrow, tomorrow) //
	fmt.Println(tomorrow)                      //

	friday := time.Weekday(5)
	fmt.Println("The day after Friday is: ", friday+1)

	saturday := time.Weekday(6)
	fmt.Println("The day after Saturday is: ", saturday+1)

	// dayAfterTomorrow := day + 2
	// fmt.Printf("%T %#v\n", dayAfterTomorrow, dayAfterTomorrow) //
	// fmt.Println(dayAfterTomorrow)                              //
}
