package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	result, err := ioutil.ReadAll(conn)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(result))
}
