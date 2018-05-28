package database

import (
	"astropay/go-web-template/logger"
	"database/sql"
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

// Generic database errors
var (
	ErrNotEnoughParams = errors.New("Not enough params for database initialization")
)

type DBAccessInterface interface {
	Init(logger logger.Logger, dbName string, params ...interface{}) error
	Get() (gdb *gorm.DB)
	DB() (db *sql.DB)
	Close() error
	Begin() *gorm.DB
	SetMaxOpenConnections(i int)
	SetMaxIdleConnections(i int)
	SetConnectionLifetime(time time.Duration)
}

var dbaImpl DBAccessInterface

func SetDBAccess(dba DBAccessInterface) {
	dbaImpl = dba
}

func Init(logger logger.Logger, dbName string, params ...interface{}) (err error) {
	if dbaImpl != nil {
		return dbaImpl.Init(logger, dbName, params)
	}

	return
}

func Get() (gdb *gorm.DB) {
	if dbaImpl != nil {
		gdb = dbaImpl.Get()
	}
	return
}

func DB() (db *sql.DB) {
	if dbaImpl != nil {
		db = dbaImpl.DB()
	}
	return
}

func Close() (err error) {
	if dbaImpl != nil {
		err = dbaImpl.Close()
	}
	return
}

func SetMaxOpenConnections(i int) {
	if dbaImpl != nil {
		dbaImpl.SetMaxOpenConnections(i)
	}
}

func SetMaxIdleConnections(i int) {
	if dbaImpl != nil {
		dbaImpl.SetMaxIdleConnections(i)
	}
}

func SetConnectionLifetime(time time.Duration) {
	if dbaImpl != nil {
		dbaImpl.SetConnectionLifetime(time)
	}
}
