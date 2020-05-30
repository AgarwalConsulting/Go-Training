package main

import (
	"fmt"
	"reflect"
)

type Greet struct {
	Greeting string
	Name     string `json:"name"`
}

func Encode(v interface{}) []byte {
	t := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	numOfFields := t.NumField()

	jString := "{"

	for i := 0; i < numOfFields; i++ {
		f := t.FieldByIndex([]int{i})
		value := rv.FieldByIndex([]int{i})

		key, ok := f.Tag.Lookup("json")

		if !ok {
			key = f.Name
		}

		if key != "-" {
			jString += fmt.Sprintf("\"%s\": \"%s\"", key, value)

			if i != numOfFields-1 {
				jString += ","
			}
		}
	}

	jString += "}"

	return []byte(jString)
}

func main() {
	g := Greet{"Howdy!", "World"}

	fmt.Printf("%s\n", Encode(g))
}
