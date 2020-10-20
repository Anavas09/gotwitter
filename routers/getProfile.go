package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Anavas09/gotwitter/db"
)

//GetProfile : Get profile data
func GetProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Please, add an ID", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "Something wrong "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(profile)
}
