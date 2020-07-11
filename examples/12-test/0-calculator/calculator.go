package calculator

func add(val ...int) int {
	sum := 0
	for i := 0; i < len(val); i++ {
		sum += val[i]
	}
	return sum
}

func Add(x, y int) int {
	return add(x, y)
}

func Mul(x, y int) int {
	return 0
}
