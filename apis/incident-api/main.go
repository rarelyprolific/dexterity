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

type question struct {
	Question   string        `json:"question"`
	AskedBy    string        `json:"askedBy"`
	AskedOn    bson.DateTime `json:"askedOn"`
	Answer     string        `json:"answer"`
	AnsweredBy string        `json:"answeredBy"`
	AnsweredOn bson.DateTime `json:"answeredOn"`
}

type incidentLog struct {
	Description string        `json:"description"`
	CreatedBy   string        `json:"createdBy"`
	CreatedOn   bson.DateTime `json:"createdOn"`
	Questions   []question    `json:"questions"`
}

type resolution struct {
	Description string        `json:"description"`
	ResolvedBy  string        `json:"resolvedBy"`
	ResolvedOn  bson.DateTime `json:"resolvedOn"`
}

type incident struct {
	ID            bson.ObjectID `bson:"_id" json:"id"`
	Summary       string        `json:"summary"`
	Status        string        `json:"status"`
	CreatedBy     string        `json:"createdBy"`
	CreatedOn     bson.DateTime `json:"createdOn"`
	LastUpdatedBy string        `json:"lastUpdatedBy"`
	LastUpdatedOn bson.DateTime `json:"lastUpdatedOn"`
	Log           []incidentLog `json:"log"`
	Resolution    resolution    `json:"resolution"`
}

func main() {
	fmt.Println("Welcome to the Dexterity Incident API")

	router := gin.Default()
	router.GET("/incidents", getIncidents)

	router.Run("localhost:8900")
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

	var incidentsFound []incident

	if err = cursor.All(ctx, &incidentsFound); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, incidentsFound)
}
