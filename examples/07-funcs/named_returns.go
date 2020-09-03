package main

import (
	"fmt"
)

func decimals(num int) (tenthsplace, onesplace int) {
	tenthsplace = num / 10
	onesplace = num - (tenthsplace * 10)

	if num == 42 {
		return 42, 0
	}

	return
}

// func doSomething() (val string,  err error) {
// 	if ... {
// 		...
// 		return val, errors.New("some helpful error message")
// 	}

// 	return
// }

func main() {
	fmt.Println(decimals(25)) // 2, 5

	fmt.Println(decimals(42)) // 42, 0
}
