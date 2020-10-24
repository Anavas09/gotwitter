package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Anavas09/gotwitter/db"
)

//ReadTweets : Read tweets from the database and return to the body
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The parameter id is required", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pag")) < 1 {
		http.Error(w, "The parameter pag is required", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("pag"))

	if err != nil {
		http.Error(w, "The parameter pag must be greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(page)

	resp, status := db.ReadTweetsFromDB(ID, pag)

	if status == false {
		http.Error(w, "Error trying read the tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}
