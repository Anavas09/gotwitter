package db

import (
	"context"
	"time"

	"github.com/Anavas09/gotwitter/models"
)

//InsertRelation : Insert a relation model to the database
func InsertRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, relation)

	if err != nil {
		return false, err
	}

	return true, nil
}
