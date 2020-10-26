package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/models"
)

//UploadBanner : Save the Banner to the server
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")

	var extension = strings.Split(handler.Filename, ".")[1]

	var filename string = "uploads/banners/" + UserID + "." + extension

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error trying upload the Banner "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error copying the Banner to the server "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Banner = UserID + "." + extension

	status, err = db.EditUser(user, UserID)

	if err != nil || status == false {
		http.Error(w, "Error trying to save the Banner in the database"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
