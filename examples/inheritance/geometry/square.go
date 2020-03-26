package main

type Square struct {
	Side float64
}

func (s *Square) Area() float64 {
	// if s == nil {
	// 	return 0
	// }
	return s.Side * s.Side
}

func (s *Square) Scale(scaleBy float64) {
	// if s == nil {
	// 	return
	// }
	s.Side = s.Side * scaleBy
}
