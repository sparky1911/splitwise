package routes

import (
	"net/http"
	"splitwise/models"
	"splitwise/utils"

	"github.com/gin-gonic/gin"
)

func createGroup(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {

		context.JSON(http.StatusUnauthorized, gin.H{"message": "token gayab"})
		return
	}
	_, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "verify nahi hua"})
		return
	}
	var group models.Group
	err = context.ShouldBindJSON(&group)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
	}
	err = group.Create()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create group"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "group created successfully"})

}
