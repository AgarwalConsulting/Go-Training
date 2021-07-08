// Describe having a separate Printer & Scanner interface
// Or

// func Save(f *os.File, doc *Document) error
// func Save(rwc io.ReadWriteCloser, doc *Document) error

package main

import (
	"fmt"
	"io"
	"os"
)

// func Save(f io.ReadWriteCloser, doc string) {
func Save(f io.Writer, doc string) {
	f.Write([]byte(doc))
}

func main() {
	f, err := os.OpenFile("/tmp/somefile.txt", os.O_CREATE|os.O_RDWR, 0744)
	defer f.Close()

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	Save(f, "Hello flipsters!")
}
