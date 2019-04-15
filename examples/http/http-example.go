package main

import "net/http"

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	msg := []byte("Hello, World!")
	w.Write(msg)
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":9000", nil)
}
