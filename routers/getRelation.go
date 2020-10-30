package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//GetRelation : Check if two user got a relationship (follow an user)
func GetRelation(w http.ResponseWriter, r *http.Request) {
	UserRelationID := r.URL.Query().Get("id")

	if len(UserRelationID) < 1 {
		http.Error(w, "The parameter id is required", http.StatusBadRequest)
		return
	}

	var relation models.Relation

	relation.UserID = UserID
	relation.UserRelationID = UserRelationID

	var resp models.ResponseCheckRelation

	status, err := db.CheckRelation(relation)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)
}
