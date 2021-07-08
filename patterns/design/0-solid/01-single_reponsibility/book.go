package main

type Book struct {
	Name      string
	Author    string
	Price     float64
	Reviewers []string
}

func (b *Book) updatePrice(newPrice float64) {
	b.Price = newPrice
}
