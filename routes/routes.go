package routes

import (
	"github.com/fcoders/api-template/controllers"
	"github.com/gin-gonic/gin"
)

// InitRoutes configures the HTTP routes defined in the API
func InitRoutes(engine *gin.Engine) {

	// health check
	engine.GET("health", controllers.Health())

}
