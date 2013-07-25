package models

import (
	"github.com/stephenalexbrowne/zoom"
)

func Initialize() error {
	config := zoom.DbConfig{
		Database:   7,
		PoolSize:   999,
		UseSockets: true,
		Address:    "/tmp/redis.sock",
	}
	zoom.InitDb(config)

	err := zoom.Register(&Person{}, "person")
	if err != nil {
		return err
	}

	return nil
}
