package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/incident-api/models"
	"github.com/rarelyprolific/dexterity/incident-api/mongoinit"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	fmt.Println("Welcome to the Dexterity Incident API")
	fmt.Println()

	// Set up the Mongo DB client
	mongoClient, err := mongoinit.CreateClient()
	if err != nil {
		log.Fatalf("Failed to initialise connection to Mongo DB: %v", err)
	}

	// Create and start the Gin router with Mongo middleware and routes
	router := CreateRouter(mongoClient)

	router.Run("0.0.0.0:8900")
}

// getIncidents get summaries for all incidents
func getIncidents(c *gin.Context) {
	client := c.MustGet("mongoClient").(*mongo.Client)
	incidentsCollection := client.Database("dexterity").Collection("incidents")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	cursor, err := incidentsCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var incidentSummaries []models.IncidentSummary

	if err = cursor.All(ctx, &incidentSummaries); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, incidentSummaries)
}

// getIncidentById gets a single incident by ID
func getIncidentById(c *gin.Context) {
	client := c.MustGet("mongoClient").(*mongo.Client)
	incidentsCollection := client.Database("dexterity").Collection("incidents")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	incidentId := c.Param("id")

	objectId, err := bson.ObjectIDFromHex(incidentId)

	if err != nil {
		// TODO: Don't expose internal errors. Handle and return RESTful responses!
		c.JSON(http.StatusBadRequest, fmt.Sprintf("invalid ObjectID: %s", err))
		return
	}

	var result models.Incident
	err = incidentsCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)

	if err != nil {
		// TODO: Don't expose internal errors. Handle and return RESTful responses!
		c.JSON(http.StatusBadRequest, fmt.Sprintf("document not found: %s", err))
		return
	}

	c.JSON(http.StatusOK, result)
}
