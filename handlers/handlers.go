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

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
