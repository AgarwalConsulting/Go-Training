package Shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Circumference() float64 {
	return 2 * r.Height * r.Width
}
