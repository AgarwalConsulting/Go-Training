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

	// tomorrow := day + 1
	// fmt.Printf("%T %#v\n", tomorrow, tomorrow) //
	// fmt.Println(tomorrow)                      //

	// dayAfterTomorrow := day + 2
	// fmt.Printf("%T %#v\n", dayAfterTomorrow, dayAfterTomorrow) //
	// fmt.Println(dayAfterTomorrow)                              //
}
