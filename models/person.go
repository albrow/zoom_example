package models

import (
	"github.com/stephenalexbrowne/zoom"
)

type Person struct {
	Name string
	Age  int
	zoom.DefaultData
}

func NewPerson(name string, age int) *Person {
	p := &Person{
		Name: name,
		Age:  age,
	}
	return p
}

func FindPersonById(id string) (*Person, error) {
	p := &Person{}
	if _, err := zoom.ScanById(id, p).Run(); err != nil {
		return p, err
	}
	return p, nil
}

func FindAllPersons() ([]*Person, error) {
	persons := make([]*Person, 0)
	if _, err := zoom.ScanAll(&persons).Run(); err != nil {
		return persons, err
	}
	return persons, nil
}
