package services

import (
	"astropay/go-web-template/database"
	"astropay/go-web-template/database/dbmysql"
	"astropay/go-web-template/logger"

	"github.com/facebookgo/inject"
)

var (
	svcContainer *ServiceContainer
)

// ServiceContainer holds the singleton instances of objects that will
// be injected as dependency in other objects
type ServiceContainer struct {
	DatabaseAccess database.DBAccessInterface `inject:""`
	Logger         logger.Logger              `inject:""`
}

// DefaultDB returns the default database access configured
func DefaultDB() database.DBAccessInterface {
	if svcContainer != nil {
		return svcContainer.DatabaseAccess
	}
	return nil
}

// DefaultLogger returns the default logger configured
func DefaultLogger() logger.Logger {
	if svcContainer != nil {
		return svcContainer.Logger
	}
	return nil
}

// Init handles the dependency injection
func Init() (err error) {

	svcContainer = new(ServiceContainer)

	// set logger
	log := logger.NewSimpleLogger()
	logger.SetLogger(log)

	// set db
	dbAccess := new(dbmysql.MySQLAccess)
	database.SetDBAccess(dbAccess)

	// instances for service container
	var graph inject.Graph
	if err = graph.Provide(
		&inject.Object{Value: svcContainer},
		&inject.Object{Value: dbAccess},
		&inject.Object{Value: log},
	); err != nil {
		return
	}

	err = graph.Populate()
	return
}

// InitMocked handles the dependency injection for testing purposes
func InitMocked() {
	// todo for testing
}
