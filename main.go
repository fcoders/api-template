package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/fcoders/api-template/services"
	"github.com/fcoders/api-template/settings"
)

var (
	settingsFile *string
	configPath   *string
)

func main() {
	loadArguments()
	appInit()
	httpServiceStart()
}

// application start
func appInit() {

	// load settings
	if err := settings.Init(*settingsFile); err != nil {
		panic(err)
	}

	// dependency manager
	err := services.Init()
	if err != nil {
		panic(fmt.Sprintf("Error initiation dependency manager: %s", err.Error()))
	}

	// configure logger
	services.DefaultLogger().EnableErrorsToSlack(settings.Get().Log.SlackEnabled)
	services.DefaultLogger().SetSlackWeebhook(settings.Get().Log.SlackWebhook)

	// init db connection
	if err := initDatabase(); err != nil {
		services.DefaultLogger().Fatal(err)
	}
}

// starts the http service
func httpServiceStart() {
	httpService := HTTPService{}
	httpService.Init()

	go httpService.Start()

	// service stops when receiving SIGINT, SIGTERM or SIGKILL signals
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	sign := <-ch

	httpService.Stop(sign.String())
	os.Exit(0)
}

// parse command line arguments
func loadArguments() {
	settingsFile = flag.String("settings", "", "Settings file")
	flag.Parse()

	// settings file location
	sFile := *settingsFile
	if sFile != "" {
		if !filepath.IsAbs(sFile) {
			sFile = filepath.Join(getAppPath(), sFile)
			settingsFile = &sFile
		}
	} else {
		// default file location is app path + settings.yml
		sFile = filepath.Join(getAppPath(), "settings.yml")
		settingsFile = &sFile
	}

}

// returns the application execution path
func getAppPath() string {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		return dir
	}

	return ""
}

func initDatabase() (err error) {

	// main database
	err = services.DefaultDB().Init(
		services.DefaultLogger(),
		settings.Get().Database.Main.Name,
		settings.Get().Database.Main.Address,
		settings.Get().Database.Main.Username,
		settings.Get().Database.Main.Password,
	)

	return
}
