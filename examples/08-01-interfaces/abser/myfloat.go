package main

type MyFloat float64

func (m MyFloat) Abs() float64 {
	if m < 0 {
		return float64(-m)
	}
	return float64(m)
}
