package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func DownloadHandler(w http.ResponseWriter, req *http.Request) {
	csvData := `
name,age,city
Gaurav,32,Chennai
	`

	w.Header().Set("Content-Type", "application/zip")

	gzip.NewWriter(w).Write([]byte(csvData))
}

func RelayHandler(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("sample.csv.zip")

	if err != nil {
		log.Println("Unable to open:", err)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/zip")
	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/serve", DownloadHandler)
	http.HandleFunc("/relay", RelayHandler)

	http.ListenAndServe(":8000", nil)
}
