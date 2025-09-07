package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/task-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetTasks gets all the tasks
func GetTasks(c *gin.Context) {
	client := c.MustGet("mongoClient").(*mongo.Client)
	tasksCollection := client.Database("dexterity").Collection("tasks")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	cursor, err := tasksCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var taskSummaries []models.TaskSummary

	if err = cursor.All(ctx, &taskSummaries); err != nil {
		log.Printf("failed to fetch tasks: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch tasks"})
		return
	}

	c.JSON(http.StatusOK, taskSummaries)
}
