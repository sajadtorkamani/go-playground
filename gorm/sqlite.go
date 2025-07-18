package main

import (
	"fmt"
	"go-playground/gorm/db/db"
)

func main() {
	db.ConnectToDatabase()

	jokes := []string{"Joke 1", "Joke 2", "Joke 3", "Joke 4"}
	numJokesAdded := 0

	for _, joke := range jokes {
		if db.JokeDoesNotExist(joke) {
			db.DB.Create(&db.Joke{Content: joke})
			numJokesAdded += 1
		}
	}

	if numJokesAdded > 0 {
		fmt.Println(
			fmt.Sprintf("✅  Added new %d jokes. We now have %d jokes in total.", numJokesAdded, db.GetJokeCount()),
		)
	} else {
		fmt.Println("No jokes added.")
	}
}
