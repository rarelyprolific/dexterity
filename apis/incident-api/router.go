package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/incident-api/mongoconnection"
	"github.com/rarelyprolific/dexterity/incident-api/routes"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// CreateRouter creates a Gin router with Mongo middleware and routes
func CreateRouter(mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.Use(mongoconnection.InjectAsMiddleware(mongoClient))

	// Set up routes to API endpoints
	router.DELETE("/incidents/:id", routes.DeleteIncidentById)
	router.GET("/incidents", routes.GetIncidents)
	router.GET("/incidents/:id", routes.GetIncidentById)

	return router
}
