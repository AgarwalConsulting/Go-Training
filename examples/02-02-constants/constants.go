package main

import (
	"fmt"
)

type status int

// const (
// 	running status = 0
// 	waiting status = 1
// 	done    status = 2
// )

const (
	running status = iota * 3
	waiting
	done
)

func main() {
	var s status

	fmt.Println(s) // Default value: 0

	fmt.Printf("%v, %T\n", running, running) // 0, Type: main.status

	fmt.Printf("%v, %T\n", done, done) // 6, Type: main.status

	// fmt.Println(time.Microsecond * 10000)

	// fmt.Println(time.Weekday(0))
}
