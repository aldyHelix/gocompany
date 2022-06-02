package routes

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")
	index := r.router.Group("/") //without V1 route API

	r.index(index)
	r.company(v1)
	r.auth(v1)
	r.addPing(v1)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
