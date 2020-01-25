package functions

func fac(num int) int {
	value := 1
	for i := 1; i <= num; i++ {
		value *= i
	}
	return value
}

func facr(num int) int {
	if num == 0 {
		return 1
	}

	return num * facr(num-1)
}
