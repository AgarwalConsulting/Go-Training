package main

import "fmt"

func change(n *[10]int) {
	n[3] = 12
}

func main() {
	a := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	change(&a)

	fmt.Println(a[3]) // 7

	for i, el := range a {
		fmt.Println(i, el)
	}

	fmt.Println(a)
}
