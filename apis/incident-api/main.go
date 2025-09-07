package main

import (
	"fmt"
	"log"

	"github.com/rarelyprolific/dexterity/incident-api/mongoconnection"
)

func main() {
	fmt.Println("Welcome to the Dexterity Incident API")
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
