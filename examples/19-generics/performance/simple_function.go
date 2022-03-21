package main

func add(x, y int) int {
	return x + y
}

type Number interface {
	int | float64
}

func addG[T Number](x, y T) T {
	return x + y
}
