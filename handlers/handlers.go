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

	//Register and Login
	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")

	//Profile
	router.HandleFunc("/getProfile", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/editProfile", middlewares.CheckDB(middlewares.ValidateJWT(routers.EditProfile))).Methods("PUT")

	//Tweets
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.SaveTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	//Avatar and Banner
	router.HandleFunc("/uploadAvatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckDB(routers.GetBanner)).Methods("GET")

	//Relations (following)
	router.HandleFunc("/addRelation", middlewares.CheckDB(middlewares.ValidateJWT(routers.AddRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/getRelation", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetRelation))).Methods("GET")

	//Users
	router.HandleFunc("/listUsers", middlewares.CheckDB(middlewares.ValidateJWT(routers.ListUsers))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
