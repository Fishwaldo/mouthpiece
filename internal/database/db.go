package database

import (
	"time"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"

	"github.com/xo/dburl"
	"github.com/spf13/viper"
)

var Conn *sql.DB
var Type string

func init() {

	viper.SetDefault("db.idle", 10)
	viper.SetDefault("db.max", 100)
	viper.SetDefault("db.lifetime", time.Hour)
	viper.SetDefault("db.connection", "sqlite3://mouthpiece.db")
}

func Start() error {
	dbConString := viper.GetString("db.connection")
	u, err := dburl.Parse(dbConString)
	if err != nil {
		return err
	}


	// cache=shared&_fk=1
	if u.Driver == "sqlite3" {
		V := u.Query()
		V.Set("cache", "shared")
		V.Set("_fk", "1")
		u.RawQuery = V.Encode()
	}

	fmt.Printf("Connecting to %s - %+v %s\n", u.Driver, u.DSN, u.Query().Encode())

	Conn, err = sql.Open(u.Driver, fmt.Sprintf("%s?%s", u.DSN, u.Query().Encode()))
	if err != nil {
		return err
	}
	Conn.SetMaxIdleConns(viper.GetInt("db.idle"))
	Conn.SetMaxOpenConns(viper.GetInt("db.max"))
	Conn.SetConnMaxLifetime(viper.GetDuration("db.lifetime"))

	Type = u.Driver

	return nil
}


