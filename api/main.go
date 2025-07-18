package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-playground/api/db"
	"go-playground/api/jokes"
	"go-playground/api/models"
	"net/http"
	"strconv"
)

func main() {
	db.ConnectToDatabase()

	seedDatabase()

	router := initialiseRouter()
	router.Run()
}

func initialiseRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apiEntry)
	router.GET("/jokes", listJokes)
	router.GET("/jokes/:jokeId", getJoke)
	router.POST("/jokes", createJoke)

	fmt.Println("🚀 Server listening on http://localhost:8080")

	return router
}

func apiEntry(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}

func listJokes(c *gin.Context) {
	c.JSON(http.StatusOK, jokes.GetJokes())
}

func getJoke(c *gin.Context) {
	jokeId, err := strconv.Atoi(c.Param("jokeId"))

	if err != nil {
		panic(err)
	}

	joke := jokes.GetJokeById(jokeId)

	if joke.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
		return
	}

	c.JSON(http.StatusOK, joke)
}

type CreateJokeRequest struct {
	Content string `json:"content" binding:"required"`
}

func createJoke(c *gin.Context) {
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

func seedDatabase() {
	JOKES := []string{"Joke 1", "Joke 2", "Joke 3", "Joke 4"}
	numJokesAdded := 0

	for _, joke := range JOKES {
		if jokes.JokeDoesNotExist(joke) {
			db.DB.Create(&models.Joke{Content: joke})
			numJokesAdded += 1
		}
	}

	if numJokesAdded > 0 {
		fmt.Println(
			fmt.Sprintf("✅  Added new %d jokes. We now have %d jokes in total.", numJokesAdded, jokes.GetJokeCount()),
		)
	} else {
		fmt.Println("No jokes added.")
	}
}
