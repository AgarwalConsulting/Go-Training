package main

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Scale(scaleBy float64) {
	s.Side = s.Side * scaleBy
}
