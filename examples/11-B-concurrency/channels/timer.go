package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.After(10 * time.Second)

	fmt.Println("Some busy work...")

	<-c
	fmt.Println("Done!")
}
