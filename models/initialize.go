package models

import (
	lib "github.com/stephenalexbrowne/models-example/model_lib"
)

func Initialize() error {

	lib.InitDb()

	err := lib.Register(Person{}, "person")
	if err != nil {
		return err
	}

	return nil

}
