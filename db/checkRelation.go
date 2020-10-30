package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckRelation : Check if two user got a relationship
func CheckRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("relations")

	condition := bson.M{
		"userid":         relation.UserID,
		"userrelationid": relation.UserRelationID,
	}

	var result models.Relation

	err := col.FindOne(ctx, condition).Decode(&result)

	fmt.Println(result)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
