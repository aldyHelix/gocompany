package routes

import (
	. "gocompany/env"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	if Env("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Use predefined header gin.PlatformXXX
	r.router.TrustedPlatform = gin.PlatformGoogleAppEngine
	// Or set your own trusted request header for another trusted proxy service
	// Don't set it to any suspect request header, it's unsafe
	r.router.TrustedPlatform = "X-CDN-IP"

	v1 := r.router.Group("/v1")
	index := r.router.Group("/") //without V1 route API

	r.index(index)
	r.company(v1)
	r.auth(v1)
	r.addPing(v1)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run(addr...)
}
