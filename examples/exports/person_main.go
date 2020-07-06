package main

import (
	"fmt"

	"exports/a/b/c/d"

	d1 "exports/d"

	p "exports/person"
)

func main() {
	fmt.Println(d.Hello) //
	fmt.Println(d1.Hello)

	gaurav := p.New("G A", 29)

	fmt.Println(gaurav) // Type: p.Person

	// p.SetAge(&p, 28)

	// fmt.Println(p)

	i := 28
	gaurav.SetAge(&i)

	refP := &gaurav // *p.Person

	fmt.Println(gaurav.GetName())
	fmt.Println(refP.GetName())
}
