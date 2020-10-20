package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//EditProfile : Edit user profile
func EditProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Wrong data "+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.EditUser(user, UserID)

	if err != nil {
		http.Error(w, "Something wrong was happend trying update the user data. Try again "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Can't update user data "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
