package main

import "fmt"

type BookPersistance struct {
	filename string
	book     Book
}

func (b *BookPersistance) save() {
	f, err := os.OpenFile(b.filename, os.O_CREATE|os.RD_WR, 0744)

	if err != nil {

	}

	fmt.Fprintln(f, b.book)
}
