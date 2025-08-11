package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, incidents)
}
