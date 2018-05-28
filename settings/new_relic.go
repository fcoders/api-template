package settings

import (
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"github.com/pkg/errors"
)

var (
	newRelicAgent newrelic.Application
)

// StartNewRelicAgent configures the New Relic agent
func StartNewRelicAgent() (err error) {

	config := newrelic.NewConfig(AppName, Get().NewRelic.LicenseKey)
	config.Enabled = Get().NewRelic.Enabled
	config.Transport = GetHTTPClient().Transport

	var errNewRelic error
	if newRelicAgent, errNewRelic = newrelic.NewApplication(config); errNewRelic != nil {
		err = errors.Wrap(errNewRelic, "Error creating New Relic application")
	}

	return
}

// NewRelicMiddleware defines the middleware for HTTP handling
func NewRelicMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		if newRelicAgent != nil {
			txn := newRelicAgent.StartTransaction(c.Request.URL.String(), c.Writer, c.Request)
			defer txn.End()
		}
		c.Next()
	}
}

// NewRelicNotifyError reports an error notifcication to NewRelic for a non web transaction
func NewRelicNotifyError(transaction string, errToNotify error) {
	if newRelicAgent != nil {
		trx := newRelicAgent.StartTransaction(transaction, nil, nil)
		defer trx.End()
		trx.NoticeError(errToNotify)
	}

	return
}
