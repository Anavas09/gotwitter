package routers

import (
	"net/http"

	"github.com/Anavas09/gotwitter/db"
)

//DeleteTweet : Delete a specific tweet
func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The parameter id is required", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweetFromDB(ID, UserID)

	if err != nil {
		http.Error(w, "Error trying delete the tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
