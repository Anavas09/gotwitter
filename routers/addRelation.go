package routers

import (
	"net/http"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//AddRelation : Add a relation between users (follow)
func AddRelation(w http.ResponseWriter, r *http.Request) {
	UserRelationID := r.URL.Query().Get("id")

	if len(UserRelationID) < 1 {
		http.Error(w, "The parameter id is required", http.StatusBadRequest)
		return
	}

	var relation models.Relation

	relation.UserID = UserID
	relation.UserRelationID = UserRelationID

	status, err := db.InsertRelation(relation)

	if err != nil {
		http.Error(w, "Something wrong was happend trying insert the relation "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Can't insert the relation "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
