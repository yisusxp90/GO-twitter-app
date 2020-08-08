package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yisusxp90/GO-twitter-app/middleware"
	"github.com/yisusxp90/GO-twitter-app/routers"
)

/* set port*/
func Handlers() {
	router := mux.NewRouter()
	/* sintaxis middleware: carpeta/Method one inside other one*/
	router.HandleFunc("/registro", middleware.CheckBD(routers.AddUser)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/view-profile", middleware.CheckBD(middleware.ValidateJWT(routers.GetProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
