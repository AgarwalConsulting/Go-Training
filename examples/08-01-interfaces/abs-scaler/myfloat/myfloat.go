package myfloat

type MyFloat float64

// func Abs(f MyFloat) float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

// Abs is a receiver function
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
