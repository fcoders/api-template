package main

import (
	"astropay/go-web-template/logger"
	"astropay/go-web-template/routes"
	"astropay/go-web-template/services"
	"astropay/go-web-template/settings"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	httpReadTimeout  = 4
	httpWriteTimeout = 8
)

// HTTPService represents the HTTP service that is initiated when the server starts
type HTTPService struct {
	engine    *gin.Engine
	waitGroup *sync.WaitGroup
}

// Init creates a new instance of the HTTP engine
func (service *HTTPService) Init() {

	// debug?
	if debug := settings.Get().App.Debug; !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	service.waitGroup = &sync.WaitGroup{}

	engine := gin.New()
	engine.Use(gin.Recovery())

	// enable New relic
	if settings.Get().NewRelic.Enabled {
		settings.StartNewRelicAgent()
		engine.Use(settings.NewRelicMiddleware())
	}

	// set gin router engine
	service.engine = engine
}

// Start starts the HTTP service
func (service *HTTPService) Start() {

	routes.InitRoutes(service.engine)
	port := fmt.Sprintf(":%v", settings.Get().App.HTTPPort)

	server := &http.Server{
		Addr:         port,
		Handler:      service.engine,
		ReadTimeout:  httpReadTimeout * time.Second,
		WriteTimeout: httpWriteTimeout * time.Second,
	}

	go server.ListenAndServe()
	service.waitGroup.Add(1)

	log := logger.GetLogger()
	log.Infof("%s service started!", settings.AppName)
	log.Infof("Version %s commit %s", settings.Version, settings.CommitHash)
}

// Stop ends the HTTP service execution and release all the resources
func (service *HTTPService) Stop(cause string) {
	log := logger.GetLogger()
	log.Infof("Shutdown requested with signal '%s'", strings.ToUpper(cause))

	// releases...
	if db := services.DefaultDB(); db != nil {
		db.Close()
	}

	log.Infof("%s service is now ready to exit, bye!", settings.AppName)
	service.waitGroup.Done()
}
