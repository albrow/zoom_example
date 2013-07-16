package models

import (
	lib "github.com/stephenalexbrowne/models-example/model_lib"
)

type Person struct {
	Name string
	Age  int
	*lib.Model
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age, &lib.Model{}}
}

func (p *Person) Save() (*Person, error) {
	// invoke the general saver
	result, err := lib.Save(*p)
	if err != nil {
		return nil, err
	}
	pers := result.(Person)
	return &pers, nil
}

func FindPersonById(id string) (*Person, error) {
	result, err := lib.FindById("person", id)
	if err != nil {
		return nil, err
	}
	p := result.(Person)
	return &p, nil
}
