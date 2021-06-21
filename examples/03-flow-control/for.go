package main

import (
	"fmt"
)

func main() {
	// // Normal for loop
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	// // for without init & post == while loop
	// i := 0
	// for i < 10 {
	// 	fmt.Println("i:", i)
	// 	i++
	// }

	// // for without any condition == forever loop
	// i := 0
	// for {
	// 	fmt.Println("i:", i)
	// 	// time.Sleep(1 * time.Second)

	// 	if i == 10 {
	// 		break
	// 	}
	// 	i++
	// }
}
