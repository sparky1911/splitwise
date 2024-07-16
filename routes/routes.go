package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/user/register", registerUser)
	server.POST("/user/login", login)
	server.POST("/groups", createGroup)
	server.POST("/memberships", createMembership)
	server.GET("/memberships", fetchAllMemberships)
	server.GET("/memberships/:id", fetchMembershipsByGroupID)
}
