package controllers

import (
	"astropay/go-web-template/settings"
	"encoding/json"
	"net/http"

	"github.com/astropay/go-tools/files"
	"github.com/gin-gonic/gin"
)

type appStatus struct {
	Status    string `json:"status"`
	Component string `json:"component"`
	Version   string `json:"version"`
	Server    string `json:"server"`
}

// Health handles the helth check endpoint logic
func Health() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseCode := http.StatusOK
		if files.Exists("/etc/app-mode/maintenance") {
			responseCode = http.StatusFound
		}

		// Load application information
		status := getStatus(responseCode)
		json, err := json.Marshal(status)
		if err == nil {
			c.Writer.Write(json)
		}

		c.Writer.WriteHeader(responseCode)
		c.Next()
	}
}

func getStatus(responseCode int) *appStatus {
	status := new(appStatus)
	status.Status = "Operational"

	if responseCode != http.StatusOK {
		status.Status = "Maintenance"
	}

	status.Component = settings.AppName
	status.Version = settings.Version
	status.Server = settings.HostName

	return status
}
