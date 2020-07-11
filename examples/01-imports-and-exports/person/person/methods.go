package person

import "time"

// SetAge is a pointer receiver on Person
func (p *Person) SetAge(age *int) {
	p.age = *age
	p.yearOfBirth = time.Now().Year() - *age
}

func (p Person) GetName() string {
	return p.name
}
