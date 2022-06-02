package controllers

import (
	db "gocompany/database"
	models "gocompany/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GET /company
//Get all company

func FindCompanies(c *gin.Context) {
	var company []models.Company

	db.Instance.Find(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

//POST /company
//Create new company

func CreateCompany(c *gin.Context) {
	//Validate input
	var input models.InsertCompany
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Create Company
	company := models.Company{CompanyName: input.CompanyName, CompanyDescription: input.CompanyDescription}
	db.Instance.Create(&company)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

//GET /company/:id
//find a company data
func FindCompany(c *gin.Context) {
	var company models.Company

	if err := db.Instance.Where("id = ?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"data": company})
}

//PATCH /company/:id
// Update a company
func UpdateCompany(c *gin.Context) {
	//get model if exist
	var company models.Company
	if err := db.Instance.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
		return
	}

	//validate input
	var input models.EditCompany
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Instance.Model(&company).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": company})
}

//DELETE company/:id
//Delete a company
func DeleteCompany(c *gin.Context) {
	//get model if exist
	var company models.Company
	if err := db.Instance.Where("id=?", c.Param("id")).First(&company).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Instance.Delete((&company))
	c.JSON(http.StatusOK, gin.H{"data": true})
}
