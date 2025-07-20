package api

import (
	"github.com/gin-gonic/gin"
	"go-playground/api/db"
	"net/http"
	"strconv"
)

func ListJokes(c *gin.Context) {
	c.JSON(http.StatusOK, db.GetJokes())
}

func GetJoke(c *gin.Context) {
	jokeId, err := strconv.Atoi(c.Param("jokeId"))

	if err != nil {
		panic(err)
	}

	joke := db.GetJokeById(jokeId)

	if joke.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, joke)
}

type CreateJokeRequest struct {
	Content string `json:"content" binding:"required"`
}

func CreateJoke(c *gin.Context) {
	// TODO: Validate that the request body is a JSON object with a `content` field
	// https://gin-gonic.com/en/docs/examples/binding-and-validation/

	var request CreateJokeRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "TODO: Create the joke"})
}
