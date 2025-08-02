package routes

import (
	"backend-go-boilerplate/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", controllers.HealthCheck)
	}
}
