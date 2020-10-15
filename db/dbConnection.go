package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN : Client to handling the database
var MongoCN = dbConnection()
var clientsOptions = options.Client().ApplyURI("mongodb://localhost:27017")

func dbConnection() *mongo.Client {
	db, err := mongo.Connect(context.TODO(), clientsOptions)
	if err != nil {
		log.Fatal(err.Error())
		return db
	}

	//Check wheter the connection was successful by pinging the MongoDB server
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Cannot connect to MongoDB: %v\n", err.Error())
	} else {
		fmt.Println("Connected to MongoDB")
	}

	return db
}

//CheckDBConnection check if the database connection still up
func CheckDBConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}

	return true
}
