package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/incident-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetIncidents get summaries for all incidents
func GetIncidents(c *gin.Context) {
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
		log.Printf("failed to fetch incidents: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch incidents"})
		return
	}

	c.JSON(http.StatusOK, incidentSummaries)
}
