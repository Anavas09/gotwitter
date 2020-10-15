package middlewares

import (
	"net/http"

	"github.com/Anavas09/gotwitter/db"
)

//CheckDB : Middleware to check if the database connection still up
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckDBConnection() == false {
			http.Error(w, "Lost Database connection", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
