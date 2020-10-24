package db

import (
	"context"
	"log"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//ReadTweetsFromDB : Read tweets from an user
func ReadTweetsFromDB(ID string, pag int64) ([]*models.ReturnTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("tweets")

	var results []*models.ReturnTweet

	condition := bson.M{"userid": ID}

	opt := options.Find()

	opt.SetLimit(20)
	opt.SetSort(bson.D{{Key: "date", Value: -1}})
	opt.SetSkip((pag - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var document models.ReturnTweet

		err := cursor.Decode(&document)

		if err != nil {
			return results, false
		}

		results = append(results, &document)
	}

	return results, true
}
