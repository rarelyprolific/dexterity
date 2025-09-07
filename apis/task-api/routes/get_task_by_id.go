package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/task-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetTaskById gets a single task by ID
func GetTaskById(c *gin.Context) {
	client := c.MustGet("mongoClient").(*mongo.Client)
	tasksCollection := client.Database("dexterity").Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	taskId := c.Param("id")

	objectId, err := bson.ObjectIDFromHex(taskId)

	if err != nil {
		// TODO: Don't expose internal errors. Handle and return RESTful responses!
		c.JSON(http.StatusBadRequest, fmt.Sprintf("invalid ObjectID: %s", err))
		return
	}

	var result models.Task
	err = tasksCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)

	if err != nil {
		// TODO: Don't expose internal errors. Handle and return RESTful responses!
		c.JSON(http.StatusBadRequest, fmt.Sprintf("document not found: %s", err))
		return
	}

	c.JSON(http.StatusOK, result)
}
