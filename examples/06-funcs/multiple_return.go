package main

import (
	"fmt"
	"os"
)

func hello() (int, string) {
	return 42, "Hello"
}

func main() {
	// fmt.Println(hello())

	// j := 0

	// j = j + 1
	// j = j - 1

	// fmt.Println(100 / j)

	// i, err := strconv.Atoi("421212aaasd")

	// if err != nil {
	// 	fmt.Println("error encountered: ", err)
	// }

	// fmt.Println("The value is", i)

	// f, err := os.Open("/tmp/somefile")
	f, err := os.OpenFile("/tmp/somefile", os.O_RDWR, 0755)

	if err != nil {
		fmt.Println("error encountered: ", err)
		return
	}

	fmt.Println("Writing to the file...")
	_, err = f.Write([]byte("hello, world!"))

	if err != nil {
		fmt.Println("error encountered: ", err)
	}

	f.Close()
}
