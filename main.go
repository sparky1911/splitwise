package main

import (
	"splitwise/db"
	"splitwise/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up the Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r)

	// Start the server
	r.Run(":8080")
}
