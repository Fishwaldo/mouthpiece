package db

import (
	_ "fmt"

	. "github.com/Fishwaldo/mouthpiece/internal/log"
	//"gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
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
