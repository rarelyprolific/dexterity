package routes

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rarelyprolific/dexterity/wiki-api/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// GetWikiPages get a list of all wiki pages
func GetWikiPages(c *gin.Context) {
	client := c.MustGet("mongoClient").(*mongo.Client)
	wikiCollection := client.Database("dexterity").Collection("wiki")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}

	cursor, err := wikiCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var wikiPages []models.Page

	if err = cursor.All(ctx, &wikiPages); err != nil {
		log.Printf("failed to fetch wiki pages: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch wiki pages"})
		return
	}

	c.JSON(http.StatusOK, wikiPages)
}
