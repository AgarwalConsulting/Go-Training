package shapes

type Square struct {
	Side float64
}

// Area is implemented on Square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s *Square) Scale(by float64) {
	s.Side = s.Side * by
}
