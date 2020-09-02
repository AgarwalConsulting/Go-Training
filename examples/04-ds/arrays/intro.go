package main

import "fmt"

func main() {
	// C
	// int arr[10];

	var arr [10]int

	fmt.Printf("%T %v\n", arr, arr)

	arr[1] = 10
	arr[4] = 12

	fmt.Println(arr)

	for i := 0; i < cap(arr); i++ {
		element := arr[i]
		fmt.Println("Index: ", i, " Element: ", element)
	}

	// for index, element := range arr {
	// 	fmt.Println("Index: ", index, " Element: ", element)
	// }
}
