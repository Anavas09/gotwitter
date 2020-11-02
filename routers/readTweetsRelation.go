package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Anavas09/gotwitter/db"
)

//ReadTweetsFollowers : Read tweets from our followers
func ReadTweetsFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pag")) < 1 {
		http.Error(w, "The parameter pag is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("pag"))

	if err != nil {
		http.Error(w, "The parameter pag must be greater than 0", http.StatusBadRequest)
		return
	}

	response, status := db.ReadTweetsFollowers(UserID, page)

	if status == false {
		http.Error(w, "Error trying read the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
