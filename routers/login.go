package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Anavas09/gotwitter/db"
	"github.com/Anavas09/gotwitter/jwt"
	"github.com/Anavas09/gotwitter/models"
)

//Login : Compare data with the database
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Email and/or password wrong "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email required", 400)
		return
	}

	document, exist := db.TryLogin(user.Email, user.Password)

	if exist == false {
		http.Error(w, "Email and/or password wrong", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Can´t generate jwt "+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(resp)

	//Save the jwt in a cookie
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
