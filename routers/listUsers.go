package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Anavas09/gotwitter/db"
)

//ListUsers : List users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "The parameter page must be a integer number greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(pageTemp)

	results, status := db.ReadAllUsers(UserID, pag, search, typeUser)

	if status == false {
		http.Error(w, "Error getting users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(results)
}
