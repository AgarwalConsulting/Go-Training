package main

import (
	"flag"
	"fmt"
)

func main() {
	var port = flag.Int("port", 8000, "port number to start the server on")

	flag.Parse()

	// log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Starting server on port:", *port)
	fmt.Printf("Option myFlag: %T, %d, %d\n", port, port, *port)
}
