package main

import (
	"fmt"
	"net/http"
)

// Writing our first "Hello, World!" API
func main() {
	// r := http.NewServeMux()

	// http.DefaultServeMux
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World!"
		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	http.ListenAndServe(":8000", nil)
}
