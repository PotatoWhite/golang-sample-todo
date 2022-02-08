package main

import (
	"log"
	"sample-todo/config/database"
	"sample-todo/expose/rest"
)

func main() {
	database.Init()
	r := rest.Init()

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
