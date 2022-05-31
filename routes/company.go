package routes

import (
	. "gocompany/controllers"

	"github.com/gin-gonic/gin"
)

func (r routes) company(rg *gin.RouterGroup) {
	company := rg.Group("/company")

	company.GET("/", FindCompanies)
	company.POST("/", CreateCompany)
	company.GET("/:id", FindCompany)
	company.PATCH("/:id", UpdateCompany)
	company.DELETE("/:id", DeleteCompany)
}
