package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mongo/routes"
)

var router *gin.Engine

func main() {
	// router = gin.Default()
	// Initialize the routes
	routes.TodoRoute()

	// Start serving the application
	router.Run(":8000")

}
