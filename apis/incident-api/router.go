package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/incident-api/mongoinit"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateRouter creates a Gin router with Mongo middleware and routes
func CreateRouter(mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(mongoinit.InjectAsMiddleware(mongoClient))

	// Set up routes to API endpoints
	router.GET("/incidents", getIncidents)
	router.GET("/incidents/:id", getIncidentById)

	return router
}
