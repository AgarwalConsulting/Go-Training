package main

import (
	"fmt"
	"time"
)

func main() {
	// // Normal for loop
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i:", i)
	// }

	// // for without init & post == while loop
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("i:", i)
	// }

	// // for without any condition == forever loop
	i := 0
	for {
		i++
		fmt.Println("i:", i)
		time.Sleep(1 * time.Second)

		if i == 10 {
			break
		}
	}
}
