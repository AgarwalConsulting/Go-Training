package main

import (
	"fmt"
	"os"
)

func closeFile(file *os.File) {
	fmt.Println("Closing file...")

	if r := recover(); r != nil {
		fmt.Println("Recovering from...", r)
	}

	file.Close()
}

func main() {
	file, err := os.Open("/tmp/non-existing.go")
	defer closeFile(file)

	if err != nil {
		fmt.Println("Encountered error: ", err)
	}

	var i, j int

	fmt.Println(i / j)

	data := make([]byte, 100)
	count, err := file.Read(data)

	if err != nil {
		fmt.Println("Encountered error: ", err)
	}

	fmt.Println("Bytes read: ", count)
}
