package routes

import (
	"net/http"
	"splitwise/models"
	"splitwise/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateMembership handles the creation of a new membership
func createMembership(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "token missing"})
		return
	}
	_, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "token verification failed"})
		return
	}

	var membership models.Membership
	err = context.ShouldBindJSON(&membership)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}
	err = membership.Create()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not create membership"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "membership created successfully"})
}

// FetchAllMemberships handles fetching all memberships
func fetchAllMemberships(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "token missing"})
		return
	}
	_, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "token verification failed"})
		return
	}

	memberships, err := models.FetchMemberships()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch memberships"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"memberships": memberships})
}

// FetchMembershipsByGroupID handles fetching memberships by group ID
func fetchMembershipsByGroupID(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "token missing"})
		return
	}
	_, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "token verification failed"})
		return
	}

	groupID, err := strconv.ParseInt(context.Param("group_id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	if groupID > 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "group ID is required"})
		return
	}

	memberships, err := models.FetchMembershipsByGroupID(groupID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch memberships for group"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"memberships": memberships})
}
