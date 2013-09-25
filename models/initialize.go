package models

import (
	"github.com/stephenalexbrowne/zoom"
)

func Initialize() error {
	zoom.Init(&zoom.Configuration{
		Database: 2,
		Network:  "unix",
		Address:  "/tmp/redis.sock",
	})

	err := zoom.Register(&Person{}, "person")
	if err != nil {
		return err
	}

	return nil
}
