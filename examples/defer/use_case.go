package main

import (
	"fmt"
	"os"
)

func closeTheFile(f *os.File) {
	fmt.Println("Closing the file...")
	f.Close()
}

func main() {
	var i, j int
	f, _ := os.Open("/tmp/file.go")

	defer closeTheFile(f)

	fmt.Println(i / j) // panic!

	fmt.Println("Is file closed?")
}
