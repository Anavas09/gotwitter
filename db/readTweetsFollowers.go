package db

import (
	"context"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ReadTweetsFollowers : Read tweets from followers
func ReadTweetsFollowers(ID string, page int) ([]models.ReturnTweetsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})

	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweets",
		},
	})

	conditions = append(conditions, bson.M{"$unwind": "$tweets"})

	conditions = append(conditions, bson.M{"$sort": bson.M{"tweets.date": -1}})

	conditions = append(conditions, bson.M{"$skip": skip})

	conditions = append(conditions, bson.M{"$limit": 20})

	var results []models.ReturnTweetsFollowers

	cursor, err := col.Aggregate(ctx, conditions)

	if err != nil {
		return results, false
	}

	err = cursor.All(ctx, &results)

	if err != nil {
		return results, false
	}

	return results, true
}
