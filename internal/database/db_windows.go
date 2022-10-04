package database

// import (
// 	"fmt"
// 	"database/sql"
// 	"database/sql/driver"

// 	mysqlite "modernc.org/sqlite"
// )

// func init() {
// 	sql.Register("sqlite", sqliteDriver{Driver: &mysqlite.Driver{}})
// }


// type sqliteDriver struct {
// 	*mysqlite.Driver
// }

// func (d sqliteDriver) Open(name string) (driver.Conn, error) {
// 	conn, err := d.Driver.Open(name)
// 	if err != nil {
// 		return conn, err
// 	}
// 	c := conn.(interface{Exec(stmt string, args []driver.Value) (driver.Result, error)})
// 	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
// 		conn.Close()
// 		return nil, fmt.Errorf("failed to enable enable foreign keys %w", err)
// 	}
// 	return conn, nil
// }