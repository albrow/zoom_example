package models

import (
	"github.com/stephenalexbrowne/zoom"
)

type Person struct {
	Id   string
	Name string
	Age  int
}

func (p *Person) GetId() string {
	return p.Id
}

func (p *Person) SetId(id string) {
	p.Id = id
}

func NewPerson(name string, age int) *Person {
	p := &Person{Name: name, Age: age}
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
