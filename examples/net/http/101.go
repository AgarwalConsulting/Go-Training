package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Place string

type Vertex struct {
	X, Y float64
}

var m map[Place]Vertex = map[Place]Vertex{
	Place("Bell Labs"): Vertex{
		40.68433, -74.39967,
	},
	Place("Google"): Vertex{
		37.42202, -122.08408,
	},
}

func handler(w http.ResponseWriter, req *http.Request) {
	jsonEncoder := json.NewEncoder(w)
	err := jsonEncoder.Encode(m)

	if err != nil {
		fmt.Println(err)
	}
	// w.Write([]byte("Hello, world!\n"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
