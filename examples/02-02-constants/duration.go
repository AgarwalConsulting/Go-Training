package main

import (
	"fmt"
	"time"
)

func main() {
	var dur time.Duration

	dur = 10_000_000_000

	fmt.Println(dur)

	fmt.Println(10 * time.Second)
	fmt.Println(30 * time.Second)
}
