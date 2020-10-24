package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Anavas09/gotwitter/middlewares"
	"github.com/Anavas09/gotwitter/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handlers : Set the port, the handler and set the server ready to listen
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/getProfile", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/editProfile", middlewares.CheckDB(middlewares.ValidateJWT(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
