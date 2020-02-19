package main

type Rectangle struct {
	Length  float64
	Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}
