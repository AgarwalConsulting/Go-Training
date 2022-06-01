package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Address struct {
	City    string `json:"city,omitempty"`
	Pincode string `json:"pincode,omitempty"`
}

type Person struct {
	Name string `json:"name"` // Struct Tags
	Age  int    `json:"yearsSinceBirth"`

	Address Address `json:"address"`
}

// json.Marshaler
// func (p Person) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "yearsSinceBirth": %d}`, p.Name, p.Age)

// 	return []byte(jsonString), nil
// }

func (p Person) String() string {
	return fmt.Sprintf("%s is %d years old.", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Gaurav", Age: 31, Address: Address{City: "Chennai"}}

	fmt.Println(p) // Gaurav is 31 years old.

	// Write person's info as json into /tmp/people.json

	f, err := os.OpenFile("/tmp/people.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()

	if err != nil {
		log.Fatalln("Unable to open file:", err) // os.Exit(1)
	}

	err = xml.NewEncoder(f).Encode(p)

	if err != nil {
		log.Println("Unable to write:", err)
		return
	}

	// fmt.Println("Written", n, "bytes successfully!")
}
