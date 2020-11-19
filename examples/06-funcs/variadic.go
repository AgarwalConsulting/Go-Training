package main

import "fmt"

// func whatever(vals ...interface{}) {
// 	fmt.Printf("%T, %d\n", vals, len(vals))

// 	for i, val := range vals {
// 		fmt.Printf("%d, %#v\n", i, val)
// 	}
// }

func sum(numbers ...int) (result int) { // ... => splat operator
	fmt.Printf("%T %v\n", numbers, numbers) // []int, {}

	// for i := 0; i < len(numbers); i++ {
	// 	result += numbers[i]
	// }
	for _, num := range numbers {
		result += num
	}

	return result
}

func main() {
	result := sum()
	fmt.Println(result)

	result = sum(10, 20)
	fmt.Println(result) // 30

	result = sum(10, 20, 30, 40, 50)
	fmt.Println(result) // 150

	// whatever("hi", true, 42, nil)
}
