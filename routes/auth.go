package routes

import (
	"gocompany/controllers"
	"gocompany/middlewares"

	"github.com/gin-gonic/gin"
)

func (r routes) auth(rg *gin.RouterGroup) {
	auth := rg.Group("/auth")

	auth.POST("/token", controllers.GenerateToken)
	auth.POST("/user/register", controllers.RegisterUser)
	secured := auth.Group("/secured").Use(middlewares.Auth())
	{
		secured.GET("/ping", controllers.Ping)
	}
}
