package models

import (
	"errors"
	"fmt"
	"github.com/stephenalexbrowne/zoom"
)

type Person struct {
	Name string
	Age  int
	*zoom.Model
}

func NewPerson(name string, age int) *Person {
	p := &Person{
		Name:  name,
		Age:   age,
		Model: new(zoom.Model),
	}
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

func FindAllPersons() ([]*Person, error) {
	results, err := zoom.FindAll("person")
	if err != nil {
		return nil, err
	}
	persons := make([]*Person, len(results))
	for i, result := range results {
		person, ok := result.(*Person)
		if !ok {
			msg := fmt.Sprintf("Coudln't type assert %+v to *Person", result)
			return nil, errors.New(msg)
		}
		persons[i] = person
	}

	return persons, nil
}
