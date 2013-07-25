package models

import (
	"github.com/stephenalexbrowne/zoom"
	"time"
)

func Initialize() error {
	config := zoom.DbConfig{
		Timeout:  10 * time.Second,
		Database: 7,
		PoolSize: 99999,
	}
	zoom.InitDb(config)

	err := zoom.Register(&Person{}, "person")
	if err != nil {
		return err
	}

	return nil
}
