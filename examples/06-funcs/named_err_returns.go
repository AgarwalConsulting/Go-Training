package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string
}

func doSomething() (Person, error) {
	f, err := os.Open("/tmp/person.json")

	if err != nil {
		return Person{}, err
	}

	var p Person

	json.NewDecoder(f).Encode(&p)

	return p, nil
}

// func doSomething() (p Person, err error) {
// 	f, err := os.Open("/tmp/person.json")

// 	if err != nil {
// 		return p, err
// 	}

// 	json.NewDecoder(f).Encode(&p)

// 	return
// }
