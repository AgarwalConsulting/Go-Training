package shapes

type Rectangle struct {
	Length, Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}
