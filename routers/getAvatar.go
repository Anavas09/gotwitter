package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Anavas09/gotwitter/db"
)

//GetAvatar : Return the Avatar to the http
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The parameter id is required", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)

	if err != nil {
		http.Error(w, "Can't find the user", http.StatusBadRequest)
		return
	}

	avatar, err := os.Open("uploads/avatars/" + profile.Avatar)

	if err != nil {
		http.Error(w, "Can't find the user", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, avatar)

	if err != nil {
		http.Error(w, "Error trying to send the Avatar to the http", http.StatusBadRequest)
	}
}
