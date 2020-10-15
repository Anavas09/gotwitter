package main

import (
	"log"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/handlers"
)

func main() {
	if !db.CheckDBConnection() {
		log.Fatalf("Cannot connect to MongoDB")
		return
	}

	handlers.Handlers()
}
