package main

import "fmt"

type Action func(int) int

func composer(actions ...Action) Action {
	return func(el int) int {
		for _, action := range actions {
			el = action(el)
		}

		return el
	}
}

func collect(elements []int, fn Action) []int {
	newElements := make([]int, len(elements))

	for i, element := range elements {
		newElements[i] = fn(element)
	}

	return newElements
}

func addBy(n int) Action {
	return func(el int) int {
		return el + n
	}
}

func main() {
	a := []int{1, 2, 3, 4}

	add2 := addBy(2)

	multiplyBy10 := func(el int) int {
		return el * 10
	}

	formula := composer(add2, multiplyBy10)

	b := collect(a, formula)

	fmt.Println(b)
}
