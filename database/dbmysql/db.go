package dbmysql

import (
	"astropay/go-web-template/database"
	"astropay/go-web-template/logger"
	"database/sql"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// MySQLAccess holds the MySQL database access implementation
type MySQLAccess struct {
	db *gorm.DB
}

// Init initialized the database access for MySQL.
// Required parameters are:
//
// - server (type string, include port number: 127.0.0.1:3306)
// - user (type string)
// - password (type string)
//
func (access *MySQLAccess) Init(logger logger.Logger, dbName string, params ...interface{}) (err error) {
	if access != nil {

		if len(params) < 3 {
			return database.ErrNotEnoughParams
		}

		server := params[0].(string)
		user := params[1].(string)
		pwd := params[2].(string)

		connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", user, pwd, server, dbName)
		fmt.Println(connString)
		access.db, err = gorm.Open("mysql", connString)
		access.db.SetLogger(logger)
		// access.db.LogMode(debug)
	}

	return
}

func (access *MySQLAccess) Get() (gdb *gorm.DB) {
	if access != nil && access.db != nil {
		gdb = access.db
	}
	return
}

func (access *MySQLAccess) DB() (db *sql.DB) {
	if access != nil && access.db != nil {
		db = access.db.DB()
	}
	return
}

func (access *MySQLAccess) Close() (err error) {
	if access != nil && access.db != nil {
		err = access.db.Close()
	}

	return
}

func (access *MySQLAccess) Begin() (trx *gorm.DB) {
	if access != nil && access.db != nil {
		trx = access.db.Begin()
	}

	return
}

func (access *MySQLAccess) SetMaxOpenConnections(i int) {
	if access != nil && access.db != nil {
		access.db.DB().SetMaxOpenConns(i)
	}
}

func (access *MySQLAccess) SetMaxIdleConnections(i int) {
	if access != nil && access.db != nil {
		access.db.DB().SetMaxIdleConns(i)
	}
}

func (access *MySQLAccess) SetConnectionLifetime(time time.Duration) {
	if access != nil && access.db != nil {
		access.db.DB().SetConnMaxLifetime(time)
	}
}
