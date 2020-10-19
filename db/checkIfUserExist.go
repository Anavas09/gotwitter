package db

import (
	"context"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckIfUserExist : Check if the user (document) its in the database
func CheckIfUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
