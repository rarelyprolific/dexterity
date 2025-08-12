package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace with your connection string if needed
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB!")

	// Access a specific database and collection
	db := client.Database("mydatabase")
	collection := db.Collection("mycollection")

	// You can now use `collection` to perform CRUD operations
	doc := map[string]interface{}{
		"name":    "The name",
		"project": "The project",
		"active":  true,
	}

	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted document with ID: %v\n", res.InsertedID)
}
