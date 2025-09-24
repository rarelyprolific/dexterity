package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// DeleteIncidentById deletes a single incident by ID
func DeleteIncidentById(c *gin.Context) {
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

	result, err := incidentsCollection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		// TODO: Don't expose internal errors. Handle and return RESTful responses!
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("internal server error: %s", err))
		return
	}

	c.JSON(http.StatusOK, result)
}
