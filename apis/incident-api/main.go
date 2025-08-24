package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	fmt.Println("Welcome to the Dexterity Incident API")

	router := gin.Default()
	router.GET("/incidents", getIncidents)

	router.Run("localhost:8900")
}

type incident struct {
	ID    int    `json:"id"`
	Title string `json:"summary"`
}

var incidents = []incident{
	{ID: 1, Title: "The website has crashed"},
	{ID: 2, Title: "User cannot log in"},
	{ID: 3, Title: "There is an intermittent crash in a back end service"},
	{ID: 4, Title: "A service won't start up due to an error"},
	{ID: 5, Title: "A service runs out of memory shortly after starting up"},
}

func getIncidents(c *gin.Context) {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace with your connection string if needed
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Access a specific database and collection
	incidents := client.Database("dexterity").Collection("incidents")

	filter := bson.M{}

	cursor, err := incidents.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var incidentsFound []bson.M

	if err = cursor.All(ctx, &incidentsFound); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, incidentsFound)
}
