package controllers

import (
	"fmt"
	db "gocompany/database"
	"gocompany/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(context *gin.Context) {
	var user models.User

	//fmt.Println(&user)

	if err := context.ShouldBindJSON(&user); err != nil {
		fmt.Print(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		context.Abort()
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		context.Abort()
		return
	}

	record := db.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error:": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"user_id":  user.ID,
		"email":    user.Email,
		"username": user.Username,
	})
}
