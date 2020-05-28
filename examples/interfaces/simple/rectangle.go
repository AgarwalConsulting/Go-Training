package main

// Rectangle implements Shape & Outline
type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r *Rectangle) Scale(scaleBy float64) {
	r.Length = r.Length * scaleBy
	r.Breadth = r.Breadth * scaleBy
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}
