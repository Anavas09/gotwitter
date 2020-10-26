package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Anavas09/gotwitter/db"
)

//GetBanner : Return the Banner to the http
func GetBanner(w http.ResponseWriter, r *http.Request) {
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

	banner, err := os.Open("uploads/banners/" + profile.Banner)

	if err != nil {
		http.Error(w, "Can't find the user", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, banner)

	if err != nil {
		http.Error(w, "Error trying to send the Banner to the http", http.StatusBadRequest)
	}
}
