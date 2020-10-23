package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//SaveTweet : Save the tweet in the database
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet

	err := json.NewDecoder(r.Body).Decode(&tweet)

	document := models.SaveTweet{
		UserID:  UserID,
		Message: tweet.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(document)

	if err != nil {
		http.Error(w, "Something wrong was happend triying insert the document to the database. Try again "+err.Error(), 400)
	}

	if status == false {
		http.Error(w, "Can't save the Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
