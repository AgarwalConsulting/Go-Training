package main

import S "github.com/Chennai-Golang/101-workshop/examples/packages/geometry/Shapes"

func main() {

	var a, b S.Shape

	a = S.Circle{1}
	b = S.Rectangle{1, 2}

	S.PrintArea(a, b)
}
