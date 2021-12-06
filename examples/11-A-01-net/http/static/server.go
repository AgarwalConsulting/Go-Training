package main

import (
	"net/http"
	"os"
)

func main() {
	fs := os.DirFS("./") // Serve all files relative to the current directory

	hfs := http.FS(fs)

	fh := http.FileServer(hfs)

	http.ListenAndServe(":8000", fh)
}
