package main

import (
	"fmt"
	"reflect"
)

type T struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
}

func main() {
	t := reflect.TypeOf(T{})

	f1, _ := t.FieldByName("f1")
	fmt.Println(f1.Tag) // f one

	f4, _ := t.FieldByName("f4")
	fmt.Println(f4.Tag) // f four and five

	f5, _ := t.FieldByName("f5")
	fmt.Println(f5.Tag) // f four and five
}
