package main

import "net/http"

func helloWorldHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":9000", nil)
}
