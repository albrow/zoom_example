package model_lib

// File contains code strictly related to Model
// and ModelInterface. No database stuff goes here.

import (
	"reflect"
)

type Model struct {
	Id string
}

type ModelInterface interface {
	SetId(string)
}

func (m *Model) SetId(id string) {
	m.Id = id
}

// maps a type to a string identifier. The string is used
// as a key in the redis database and for the *ById methods
var typeToName map[reflect.Type]string = make(map[reflect.Type]string)
var nameToType map[string]reflect.Type = make(map[string]reflect.Type)

// adds a model to the map of registered models
// Both name and typeOf(m) must be unique, i.e.
// not already registered
func Register(m interface{}, name string) error {
	if alreadyRegisteredName(name) {
		return NewNameAlreadyRegisteredError(name)
	}
	typ := reflect.TypeOf(m)
	if alreadyRegisteredType(typ) {
		return NewTypeAlreadyRegisteredError(typ)
	}
	typeToName[typ] = name
	nameToType[name] = typ
	return nil
}

// returns true iff the model name has already been registered
func alreadyRegisteredName(n string) bool {
	for _, name := range typeToName {
		if name == n {
			return true
		}
	}
	return false
}

// returns true iff the model type has already been registered
func alreadyRegisteredType(t reflect.Type) bool {
	for typ, _ := range typeToName {
		if typ == t {
			return true
		}
	}
	return false
}
