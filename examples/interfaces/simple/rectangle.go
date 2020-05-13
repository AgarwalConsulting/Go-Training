package main

// Rectangle respresents a 2-D rectangle
type Rectangle struct {
	Length  float64
	Breadth float64
}

// Area computes the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Breadth * r.Length
}

func (r *Rectangle) Scale(scaleBy float64) {
	r.Length = r.Length * scaleBy
	r.Breadth = r.Breadth * scaleBy
}

// Perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Breadth + r.Length)
}
