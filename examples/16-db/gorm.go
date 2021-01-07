package main

import (

	// "github.com/jmoiron/sqlx"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "github.com/mattn/go-sqlite3"
	// _ "github.com/lib/pq"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://localhost/template1?sslmode=disable"), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect:", err)
	}

	db.AutoMigrate(&Person{})

	people := []Person{}

	db.Where("name = ?", "Prajwal").Find(&people)

	for _, person := range people {
		fmt.Println(person)
	}

	person := Person{Name: "Prajwal", Age: 29}

	db.Create(&person)
}
