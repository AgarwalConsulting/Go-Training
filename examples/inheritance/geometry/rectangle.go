package main

type Rectangle struct {
	Length, Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func (r *Rectangle) Scale(scaleBy float64) {
	r.Length = r.Length * scaleBy
	r.Breadth = r.Breadth * scaleBy
}
