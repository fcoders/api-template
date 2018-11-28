package services

import (
	"fmt"
	"os"
	"path"

	"github.com/facebookgo/inject"
	"github.com/fcoders/api-template/common"
	"github.com/fcoders/api-template/database"
	"github.com/fcoders/api-template/database/mysql"
	"github.com/fcoders/api-template/settings"
	"github.com/fcoders/logger"
)

var (
	svcContainer *ServiceContainer
)

// ServiceContainer holds the singleton instances of objects that will
// be injected as dependency in other objects
type ServiceContainer struct {
	DatabaseAccess database.DBAccessInterface `inject:""`
	Storage        database.Storage           `inject:""`
	Log            *logger.Logger
}

// DefaultDB returns the default database access configured
func DefaultDB() database.DBAccessInterface {
	if svcContainer != nil {
		return svcContainer.DatabaseAccess
	}
	return nil
}

// DefaultLogger returns the default logger configured
func DefaultLogger() *logger.Logger {
	if svcContainer != nil {
		return svcContainer.Log
	}
	return nil
}

// DefaultStorage returns the configured db Storage instance
func DefaultStorage() database.Storage {
	if svcContainer != nil {
		return svcContainer.Storage
	}
	return nil
}

// Init handles the dependency injection
func Init() (err error) {

	svcContainer = new(ServiceContainer)

	// ====== init logger ======
	logLocation := settings.Get().Log.Location
	if !path.IsAbs(logLocation) {
		logLocation = path.Join(common.GetAppPath(), logLocation)
	}

	// check if folder exists
	if !common.Exists(logLocation) {
		if errMakeDir := os.MkdirAll(logLocation, 0777); errMakeDir != nil {
			err = fmt.Errorf("Error creating log location path: %s", errMakeDir)
			return
		}
	}

	lf, errOpenFile := os.OpenFile(path.Join(logLocation, common.DefaultLogFileName), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if errOpenFile != nil {
		err = fmt.Errorf("Failed to open log file: %v", errOpenFile)
		return
	}

	svcContainer.Log = logger.Init(settings.AppName, settings.Get().Log.Console, settings.Get().Log.Syslog, lf)

	// ====== init db ======
	dbAccess := new(mysql.MySQLAccess)
	database.SetDBAccess(dbAccess)

	// instances for service container
	var graph inject.Graph
	if err = graph.Provide(
		&inject.Object{Value: svcContainer},
		&inject.Object{Value: dbAccess},
		&inject.Object{Value: new(mysql.Storage)},
	); err != nil {
		return
	}

	err = graph.Populate()
	return
}
