package main

import "net/http"

func helloWorldHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":9000", nil)
}
