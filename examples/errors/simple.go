package main

import (
	"errors"
	"fmt"
	"time"
)

func Double(i interface{}) (interface{}, error) {
	switch v := i.(type) {
	case int:
		return v * 2, nil
	case string:
		return v + v, nil
	default:
		errorMsg := fmt.Sprintf("Cannot double unknown type: %#v", i)
		return nil, errors.New(errorMsg)
	}
}

func main() {
	fmt.Println(Double(21))
	fmt.Println(Double("Hi "))
	fmt.Println(Double(time.Millisecond))
}
