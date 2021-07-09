package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

// Singleton: singletonDatabase; define methods on it; Read from file using `sync.Once`
// Singleton & Dependency Inversion: Don't directly call GetSingletonDatabase()

type Country struct {
	Name    string
	Capital string
}

type CountryRepository interface {
	GetCountries() []Country
	GetCapital(name string) (string, bool)
}

type Countries struct {
	countries []Country
}

func (cs *Countries) GetCountries() []Country {
	return cs.countries
}

func (cs *Countries) GetCapital(name string) (string, bool) {
	for _, c := range cs.countries {
		if c.Name == name {
			return c.Capital, true
		}
	}
	return "", false
}

var cs = Countries{}
var once sync.Once

func NewCountries(fileName string) CountryRepository {
	loadCountries := func() {
		cs.countries = make([]Country, 0)
		fmt.Println("Loading countries... from", fileName)

		// Open file
		f, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		// Read file
		r := csv.NewReader(f)

		r.Read() // skip header

		for {
			row, err := r.Read()
			if err == io.EOF {
				break
			}
			c := Country{
				Name:    row[0],
				Capital: row[1],
			}

			cs.countries = append(cs.countries, c)
		}
	}

	once.Do(loadCountries)

	return &cs
}

func main() {
	cs := NewCountries("country.csv")

	fmt.Println(cs.GetCountries())

	cs1 := NewCountries("country.csv")

	fmt.Println(cs1.GetCountries())

	fmt.Println(cs.GetCapital("Argentina"))
}
