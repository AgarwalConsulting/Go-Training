package main

import (
	"fmt"
	"strconv"
)

func main() {
	var err error // Default: <nil>

	i, err := strconv.Atoi("123ax")

	if err != nil {
		fmt.Printf("%T %v\n", err, err)
	}

	fmt.Println("Num:", i)

	// ---

	// f, err := os.Open("/tmp/restricted")
	// defer f.Close()

	// if err != nil {
	// 	fmt.Printf("%T %#v %v\n", err, err, err)
	// 	os.Exit(1)
	// } else {
	// 	fmt.Println("Found the file!")
	// }

	// // ---

	// err = errors.New("some failure")
	// if err != nil {
	// 	fmt.Printf("%T %v\n", err, err)
	// }

	// ---

	fmt.Println("Program will continue!")
}
