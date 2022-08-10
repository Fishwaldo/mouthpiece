package db

import (
	_ "fmt"
		"gorm.io/gorm"
		"gorm.io/driver/sqlite"
		. "github.com/Fishwaldo/mouthpiece/internal/log"
	)

var Db *gorm.DB

func InitializeDB() {
	var err error
	Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Log.Info("Database Initialized")
}