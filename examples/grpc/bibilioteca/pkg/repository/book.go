package repository

// SequenceGenerator function returns a generator which increments & returns the next value
func SequenceGenerator(initValue int) func() int {
	i := initValue

	return func() int {
		i++
		return i
	}
}
