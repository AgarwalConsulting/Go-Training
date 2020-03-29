package entity

// Book represents a single book
type Book struct {
	ID
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
}
