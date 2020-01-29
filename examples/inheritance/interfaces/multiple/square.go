package main

// Square conforms to Shape
type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}
