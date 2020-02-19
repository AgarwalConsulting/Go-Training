package main

type Square float64

func (s *Square) Area() float64 {
	return float64(*s * *s)
}

func (s *Square) Scale() {
	*s = *s * 10
}
