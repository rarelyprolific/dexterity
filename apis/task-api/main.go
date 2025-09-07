package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/task-api/models"
	"github.com/rarelyprolific/dexterity/task-api/mongoconnection"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	fmt.Println("Welcome to the Dexterity Task API")
	fmt.Println()

	// Set up the Mongo DB client
	mongoClient, err := mongoconnection.CreateClient()
	if err != nil {
		log.Fatalf("Failed to initialise connection to Mongo DB: %v", err)
	}

	// Create and start the Gin router with Mongo middleware and routes
	router := CreateRouter(mongoClient)

	router.Run("0.0.0.0:8900")
}

// getTasks gets all the tasks
func getTasks(c *gin.Context) {
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
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, taskSummaries)
}

// getTaskById gets a single task by ID
func getTaskById(c *gin.Context) {
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
