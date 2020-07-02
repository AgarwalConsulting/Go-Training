package shapes

type Rectangle struct {
	Length  float64
	Breadth float64
}

// Area is a Receiver function
func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

func (r *Rectangle) Scale(by float64) {
	r.Breadth = r.Breadth * by
	r.Length = r.Length * by
}
