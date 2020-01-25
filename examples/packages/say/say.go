package say

import "fmt"

const World string = "World!"

type S struct {
	f   int
	Age int
}

func NewS() S {
	return S{f: 42}
}

func Say() {
	fmt.Println("Hello!")
}
