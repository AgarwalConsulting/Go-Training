package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, _ := plugin.Open("./plugin.so")

	v, _ := p.Lookup("V")

	fmt.Printf("%#v\n", v)

	*v.(*int) = 42

	f, _ := p.Lookup("F")

	fmt.Printf("%#v\n", f)

	f.(func())()
}
