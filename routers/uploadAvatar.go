package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//UploadAvatar : Save the Avatar to the server
func UploadAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]

	var filename string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error trying upload the Avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error copying the Avatar to the server "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserID + "." + extension

	status, err = db.EditUser(user, UserID)

	if err != nil || status == false {
		http.Error(w, "Error trying to save the Avatar in the database"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
