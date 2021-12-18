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

	f, err := p.Lookup("F")

	if err != nil {
		fmt.Println("Unable to locate F:", err)
		return
	}

	fmt.Printf("%#v\n", f)

	f.(func())()
}
