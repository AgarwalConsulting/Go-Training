package main

import (
	"fmt"
	"time"
)

func Double(i interface{}) (interface{}, error) {
	// <var-of-interface>.(<type>)
	v, ok := i.(int)

	if ok {
		return 2 * v, nil
	}

	// fmt.Println(i.(int))

	switch v := i.(type) {
	case string:
		return v + v, nil
	case time.Duration:
		return v * 2, nil
	default:
		return nil, fmt.Errorf("unsupported type: %T , value: %v", i, i)
		// return nil, errors.New(errorMsg)
	}
}

func main() {
	fmt.Println(Double(21))
	fmt.Println(Double("Hi "))
	fmt.Println(Double(time.Millisecond))
	fmt.Println(Double(true))
}
