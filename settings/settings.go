// Package settings contains all the logic required to get configuration values from files or
// local environment variables
package settings

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/fcoders/api-template/common"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Settings is the structure used to hold configuration from settings.yml
type Settings struct {
	App struct {
		Debug    bool `yaml:"debug"`
		HTTPPort int  `yaml:"http_port"`
		LogLevel int  `yaml:"log_level"`
	} `yaml:"app"`
	Log struct {
		Location     string `yaml:"location"`
		Console      bool   `yaml:"console"`
		Syslog       bool   `yaml:"syslog"`
		SlackEnabled bool   `yaml:"slack_enabled"`
		SlackWebhook string `yaml:"slack_webhook"`
	} `yaml:"log"`
	Database struct {
		Main struct {
			Name               string `yaml:"name"`
			Address            string `yaml:"address"`
			Username           string `yaml:"username"`
			Password           string `yaml:"password"`
			MaxOpenConnections int    `yaml:"max_open_connections"`
			MaxIdleConnections int    `yaml:"max_idle_connections"`
			MaxLifetime        int    `yaml:"max_lifetime"`
		} `yaml:"main"`
	} `yaml:"database"`
	Proxy struct {
		Enabled bool   `yaml:"enabled"`
		Address string `yaml:"address"`
	} `yaml:"proxy"`
}

var cfg *Settings

// Init loads application settings in the indicated file
func Init(file string) (err error) {
	if common.Exists(file) {
		err = loadSettingsFromFile(file)
	} else {
		return errors.Errorf("File '%s' not found", file)
	}

	// get current host name
	HostName, _ = os.Hostname()

	return
}

// Get returns the configuration loaded
func Get() (s *Settings) {
	if cfg != nil {
		s = cfg
	}
	return
}

// GetHTTPClient returns a HTTP client with or without proxy configured
func GetHTTPClient() (client *http.Client) {
	if Get().Proxy.Enabled {
		proxyURL, _ := url.Parse(Get().Proxy.Address)
		client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	} else {
		client = http.DefaultClient
	}
	return
}

// load settings from yaml file
func loadSettingsFromFile(file string) (err error) {
	cfg = new(Settings)

	fileContent, err := ioutil.ReadFile(file)
	if err != nil {
		return errors.Wrap(err, "Error reading configuration file")
	}

	if err := yaml.Unmarshal(fileContent, cfg); err != nil {
		return errors.Wrap(err, "Error parsing configuration file")
	}

	return
}
