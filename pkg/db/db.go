package db

import (
	"context"
	_ "fmt"

	"github.com/Fishwaldo/mouthpiece/pkg/log"

	"gorm.io/gorm"
	//ll "gorm.io/gorm/logger"

	"github.com/vchitai/logrgorm2"
)

var Db *gorm.DB

func Initialize(db gorm.Dialector) *gorm.DB {
	var err error
	logger := logrgorm2.New(log.Log.WithName("Database"))
	logger = logger.IgnoreRecordNotFoundError(true)
	Db, err = gorm.Open(db, &gorm.Config{Logger: logger})
	if err != nil {
		panic(err)
	}
	//Db.Logger = logger
	//Db.Logger.LogMode(ll.Info)
	Db.Logger.Error(context.Background(), "Database Connection Established")

	log.Log.Info("Database Initialized")
	return Db
}
