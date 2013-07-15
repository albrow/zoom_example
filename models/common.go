package models

import (
	"code.google.com/p/tcgl/redis"
	"fmt"
	"time"
)

var db *redis.Database

func Db() *redis.Database {
	if db != nil {
		return db
	} else {
		fmt.Println("connecting to database...")
		config := redis.Configuration{
			Timeout:  time.Second * 10,
			Database: 7,
			PoolSize: 99999,
		}

		db := redis.Connect(config)
		return db
	}
}

func CloseDb() {
	if db != nil {
		fmt.Println("closing database...")
		db.Close()
	}
}
