package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//FindProfile : Find a profile in the database
func FindProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := col.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		fmt.Println("Can't find the user " + err.Error())

		return profile, err
	}

	return profile, nil
}
