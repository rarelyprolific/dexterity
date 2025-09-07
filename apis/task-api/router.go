package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/task-api/mongoconnection"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateRouter creates a Gin router with Mongo middleware and routes
func CreateRouter(mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(mongoconnection.InjectAsMiddleware(mongoClient))

	// Set up routes to API endpoints
	router.GET("/tasks", getTasks)
	router.GET("/tasks/:id", getTaskById)

	return router
}
