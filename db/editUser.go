package db

import (
	"context"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//EditUser : Edit user data in the database
func EditUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("users")

	document := make(map[string]interface{})

	if len(user.Name) > 0 {
		document["name"] = user.Name
	}

	if len(user.Lastname) > 0 {
		document["lastname"] = user.Lastname
	}

	if len(user.Avatar) > 0 {
		document["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		document["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		document["biography"] = user.Biography
	}

	if len(user.Location) > 0 {
		document["location"] = user.Location
	}

	if len(user.Website) > 0 {
		document["website"] = user.Website
	}

	document["birthday"] = user.Birthday

	updateString := bson.M{"$set": document}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
