package functions

// 5! => 1 * 2 * 3 * 4 * 5
func fac(num int) int {
	value := 1
	for i := 1; i <= num; i++ {
		value *= i
	}
	return value
}

// var memoize = map[int]int{}

// 5! => 5 * 4!
// 5! => 5 * 4 * 3!
// 5! => 5 * 4 * 3 * 2!
// 5! => 5 * 4 * 3 * 2 * 1!
// 5! => 5 * 4 * 3 * 2 * 1
func facr(num int) int {
	if num == 0 {
		return 1
	}

	// if v, ok := memoize[num]; ok {
	// 	return v
	// }

	v := num * facr(num-1)
	// memoize[num] = v

	return v
}
