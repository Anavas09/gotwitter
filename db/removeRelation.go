package db

import (
	"context"
	"time"

	"github.com/Anavas09/gotwitter/models"
)

//RemoveRelation : Remove a relation model from the database
func RemoveRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, relation)

	if err != nil {
		return false, err
	}

	return true, nil
}
