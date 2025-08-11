package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Welcome to the Dexterity Question API")

	router := gin.Default()
	router.GET("/questions", getQuestions)

	router.Run("localhost:8910")
}

type question struct {
	ID           int    `json:"id"`
	QuestionText string `json:"question"`
}

var questions = []question{
	{ID: 1, QuestionText: "How does this work?"},
	{ID: 2, QuestionText: "What should be done here?"},
	{ID: 3, QuestionText: "Who needs to know about this?"},
	{ID: 4, QuestionText: "What does this UI need to look like?"},
	{ID: 5, QuestionText: "How does the user authenticate?"},
}

func getQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, questions)
}
