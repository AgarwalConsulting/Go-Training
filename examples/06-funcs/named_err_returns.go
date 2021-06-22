package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
}

func doSomething() (Person, error) {
	f, err := os.Open("/tmp/person.json")
	defer f.Close()

	if err != nil {
		return Person{}, err
	}

	var p Person

	json.NewDecoder(f).Decode(&p)

	return p, nil
}

// func doSomething() (p Person, err error) {
// 	f, err := os.Open("/tmp/person.json")
// 	defer f.Close()

// 	if err != nil {
// 		return
// 	}

// 	json.NewDecoder(f).Decode(&p)

// 	return
// }

func main() {
	fmt.Println(doSomething())
}
