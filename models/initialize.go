package models

import (
	"github.com/stephenalexbrowne/zoom"
)

func Initialize() error {
	zoom.Init()

	err := zoom.Register(&Person{}, "person")
	if err != nil {
		return err
	}

	return nil
}
