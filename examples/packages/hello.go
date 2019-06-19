package main

import (
	"fmt"

	"github.com/Chennai-Golang/101-workshop/examples/packages/say"
)

func main() {
	say.Say()
	bark()

	fmt.Println(say.World)
}
