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

type Question struct {
	Question   string        `json:"question"`
	AskedBy    string        `json:"askedBy"`
	AskedOn    bson.DateTime `json:"askedOn"`
	Answer     string        `json:"answer"`
	AnsweredBy string        `json:"answeredBy"`
	AnsweredOn bson.DateTime `json:"answeredOn"`
}

type IncidentLog struct {
	Description string        `json:"description"`
	CreatedBy   string        `json:"createdBy"`
	CreatedOn   bson.DateTime `json:"createdOn"`
	Questions   []Question    `json:"questions"`
}

type Resolution struct {
	Description string        `json:"description"`
	ResolvedBy  string        `json:"resolvedBy"`
	ResolvedOn  bson.DateTime `json:"resolvedOn"`
}

type Incident struct {
	ID            bson.ObjectID `bson:"_id" json:"id"`
	Summary       string        `json:"summary"`
	Status        string        `json:"status"`
	CreatedBy     string        `json:"createdBy"`
	CreatedOn     bson.DateTime `json:"createdOn"`
	LastUpdatedBy string        `json:"lastUpdatedBy"`
	LastUpdatedOn bson.DateTime `json:"lastUpdatedOn"`
	Log           []IncidentLog `json:"log"`
	Resolution    Resolution    `json:"resolution"`
}

func main() {
	fmt.Println("Welcome to the Dexterity Incident API")
	fmt.Println()

	router := gin.Default()

	// Set up Mongo DB connection and inject into Gin context as middleware.
	client, err := initialiseMongoDbClient("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to initialise connection to Mongo DB: %v", err)
	}

	router.Use(mongoMiddleware(client))

	// Set up routes to API endpoints
	router.GET("/incidents", getIncidents)
	router.GET("/incidents/:id", getIncidentById)

	router.Run("localhost:8900")
}

// initialiseMongoDbClient sets up the connection to Mongo DB
func initialiseMongoDbClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().
		ApplyURI(uri).
		// Set connection pool size
		SetMaxPoolSize(20)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, err
	}

	// Ping to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

// mongoMiddleware adds the Mongo DB client to Gin as middleware
func mongoMiddleware(client *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mongoClient", client)
		c.Next()
	}
}

// getIncidents gets all the incidents
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

	var incidents []Incident

	if err = cursor.All(ctx, &incidents); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, incidents)
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

	var result Incident
	err = incidentsCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&result)

	if err != nil {
		// TODO: Don't expose internal errors. Handle and return RESTful responses!
		c.JSON(http.StatusBadRequest, fmt.Sprintf("document not found: %s", err))
		return
	}

	c.JSON(http.StatusOK, result)
}
