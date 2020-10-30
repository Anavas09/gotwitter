package db

import (
	"context"
	"fmt"
	"time"

	"github.com/Anavas09/gotwitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*ReadAllUsers : Got users from database. If a "R" is sended as a parameter,
got only the users wich got a relationship*/
func ReadAllUsers(ID string, page int64, search string, searchType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("gotwitterdata")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()

	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, add bool

	for cursor.Next(ctx) {
		var user models.User

		err := cursor.Decode(&user)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var relation models.Relation

		relation.UserID = ID
		relation.UserRelationID = user.ID.Hex()

		add = false

		found, err = CheckRelation(relation)

		if searchType == "new" && found == false {
			add = true
		}

		if searchType == "following" && found == true {
			add = true
		}

		if relation.UserRelationID == ID {
			add = false
		}

		if add == true {
			user.Banner = ""
			user.Biography = ""
			user.Email = ""
			user.Location = ""
			user.Password = ""
			user.Website = ""

			results = append(results, &user)
		}
	}

	err = cursor.Err()

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)

	return results, true
}
