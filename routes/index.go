package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r routes) index(rg *gin.RouterGroup) {
	ping := rg.Group("")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
}
