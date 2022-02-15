package vertex

import "math"

type Vertex struct {
	X, Y float64
	// data []float64
}

// Normal Function
// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// Abs is a value receiver
func (v Vertex) Abs() float64 {
	result := math.Sqrt(v.X*v.X + v.Y*v.Y)

	return result
}

// Abs is a pointer receiver
// func (v *Vertex) Abs() float64 {
// 	if v == nil {
// 		return 0
// 	}

// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

// Scale is a pointer receiver
func (v *Vertex) Scale(by float64) {
	if v == nil {
		return
	}
	v.X *= by
	// (*v).X *= by
	v.Y *= by
	// (*v).Y *= by
}
