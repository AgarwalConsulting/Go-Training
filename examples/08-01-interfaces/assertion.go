package main

import (
	"errors"
	"fmt"
	"time"
)

func Double(i interface{}) (interface{}, error) {
	// <var-of-interface>.(<type>)
	fmt.Printf("%T, %v\n", i, i) // string, "Hi "

	// if v, ok := i.(int); ok {
	// 	return 2 * v, nil
	// }

	switch v := i.(type) {
	case int:
		return v * 2, nil
	case string:
		return v + v, nil
	case time.Duration:
		return v * 2, nil
	default:
		return nil, errors.New("unsupported operation")
	}
}

func main() {
	fmt.Println(Double(21))
	fmt.Println(Double("Hi "))
	fmt.Println(Double(time.Millisecond))
	fmt.Println(Double(true))
}
