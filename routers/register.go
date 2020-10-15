package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//Register Add a new user to the database
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error in the received data "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email required", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", 400)
		return
	}

	_, founded, _ := db.CheckIfUserExist(user.Email)

	if founded {
		http.Error(w, "An user is registrated with the same email", 400)
		return
	}

	_, status, err := db.AddUser(user)

	if err != nil {
		http.Error(w, "Error adding a new user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Can´t add an user", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
