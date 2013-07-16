package model_lib

// File contains code strictly related to the database

import (
	"code.google.com/p/tcgl/redis"
	"fmt"
	"time"
)

var db *redis.Database

// initializes and returns the database instance
func InitDb() *redis.Database {
	if db != nil {
		return db
	} else {
		fmt.Println("connecting to database...")
		config := redis.Configuration{
			Timeout:  10 * time.Second,
			Database: 7,
			PoolSize: 99999,
		}

		db = redis.Connect(config)
		return db
	}
}

// closes the database
func CloseDb() {
	if db != nil {
		fmt.Println("closing database...")
		db.Close()
	}
}
