package routes

import (
	"gocompany/controllers"

	"github.com/gin-gonic/gin"
)

func (r routes) company(rg *gin.RouterGroup) {
	company := rg.Group("/company")

	company.GET("/", controllers.FindCompanies)
	company.POST("/", controllers.CreateCompany)
	company.GET("/:id", controllers.FindCompany)
	company.PATCH("/:id", controllers.UpdateCompany)
	company.DELETE("/:id", controllers.DeleteCompany)
}
