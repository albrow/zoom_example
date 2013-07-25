package models

import (
	"github.com/stephenalexbrowne/zoom"
)

type Person struct {
	Name string
	Age  int
	*zoom.Model
}

func NewPerson(name string, age int) *Person {
	p := &Person{Name: name, Age: age}
	p.Model = zoom.NewModelFor(p)
	return p
}

func FindPersonById(id string) (*Person, error) {
	result, err := zoom.FindById("person", id)
	if err != nil {
		return nil, err
	}
	p := result.(*Person)
	return p, nil
}
