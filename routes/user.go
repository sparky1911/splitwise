package routes

import (
	"net/http"
	"splitwise/models"
	"splitwise/utils"

	"github.com/gin-gonic/gin"
)

func registerUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user created succesfully"})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	err = user.ValidateUser()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not verify user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user logged in", "token": token})

}
