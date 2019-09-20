package shapes

type Square struct {
	Side float64
}

func (c *Square) Area() float64 {
	return c.Side * c.Side
}

func (sq *Square) Double() {
	sq.Side = sq.Side * 2
}
