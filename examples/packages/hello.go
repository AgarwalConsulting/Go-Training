package main

import (
	"fmt"
	"reflect"

	"github.com/Chennai-Golang/101-workshop/examples/packages/say"
)

func main() {
	var s say.S

	s = say.NewS()

	fmt.Println(s)

	t := reflect.ValueOf(s)

	val := t.FieldByName("Age")

	fmt.Println(val)

	val.SetInt(0)

	fmt.Println(s)
}
