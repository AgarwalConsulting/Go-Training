package main

import (
	"os"
	"fmt"
)

var port string

func init() {
	var ok bool
	port, ok = os.LookupEnv("PORT")

	if !ok {
		port = "8000"
	}

}

func main() {
	fmt.Println("Starting server on port:", port)
}
