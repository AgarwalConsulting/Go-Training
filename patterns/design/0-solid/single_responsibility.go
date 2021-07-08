package main

import (
	"fmt"
	"os"
)

// Design a function which can update the price of a book and save it's details
// Design a Book with saving logic

type Book struct {
	Name      string
	Author    string
	Price     float64
	Reviewers []string
}

func (b *Book) updatePrice(newPrice float64) {
	b.Price = newPrice
}

// func (b *Book) updatePriceAndSave(newPrice float64, filename string) {
// 	b.Price = newPrice

// 	b.save(filename)
// }

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

func main() {
	b := Book{"Art of War", "Sun Tzu", 10.12, []string{"Abhinav", "Ayushi"}}

	bp := BookPersistance{"/tmp/art_of_war.txt", b}

	bp.save()
}
