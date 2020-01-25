package main

import (
	"fmt"
	"time"

	validate "gopkg.in/go-playground/validator.v9"
)

type Inner struct {
	StartDate time.Time
}

type Outer struct {
	InnerStructField *Inner
	Age              int       `validate:"max=10"`
	Location         string    `validate:"required"`
	CreatedAt        time.Time `validate:"ltecsfield=InnerStructField.StartDate"`
}

func main() {
	now := time.Now()

	inner := &Inner{
		StartDate: now,
	}

	outer := &Outer{
		InnerStructField: inner,
		Age:              1000,
		CreatedAt:        now,
	}

	validator := validate.New()

	errs := validator.Struct(outer)

	fmt.Println(errs)
}
