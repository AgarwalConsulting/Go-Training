package main

import (
	"fmt"
	"time"

	validate "gopkg.in/go-playground/validator.v10"
)

type Birth struct {
	Date  time.Time
	Place string
}

type StudentEnrollment struct {
	Name       string
	Birth      *Birth
	Age        int       `validate:"min=3,max=5"`
	Parents    string    `validate:"required"`
	EnrolledAt time.Time `validate:"ltecsfield=Birth.Date"`
}

func main() {
	now := time.Now()

	birth := &Birth{
		Date:  now,
		Place: "Chennai",
	}

	s := &StudentEnrollment{
		Name:       "G A",
		Birth:      birth,
		Age:        2,
		EnrolledAt: now,
	}

	validator := validate.New()

	errs := validator.Struct(s)

	fmt.Println(errs)
}
