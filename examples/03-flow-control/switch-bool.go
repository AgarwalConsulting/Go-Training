package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	os := runtime.GOOS
	switch {
	case time.Now().Weekday() == 3:
		fmt.Println("Argh! Weekend is so far!")
	case os == "linux":
		fmt.Println("Linus rocks!")
	case os == "darwin":
		fmt.Println("Steve rocks!")
	default:
		fmt.Println("This OS sucks!")
	}
}
