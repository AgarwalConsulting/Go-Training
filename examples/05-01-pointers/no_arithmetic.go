package main

import "fmt"

func main() {
	var a = [4]int{2, 3, 5, 7}

	var p *int
	p = &a[1]

	p1 := p + 8 // no pointer arithematic

	fmt.Println(*p1)
}
