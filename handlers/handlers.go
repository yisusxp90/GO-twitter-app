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
	router.HandleFunc("/registro", middleware.CheckBD(routers.AddUser)).Methods(http.MethodPost)
	router.HandleFunc("/login", middleware.CheckBD(routers.Login)).Methods(http.MethodPost)
	router.HandleFunc("/view-profile", middleware.CheckBD(middleware.ValidateJWT(routers.GetProfile))).Methods(http.MethodGet)
	router.HandleFunc("/update-profile", middleware.CheckBD(middleware.ValidateJWT(routers.UpdateProfile))).Methods(http.MethodPut)
	router.HandleFunc("/tweet", middleware.CheckBD(middleware.ValidateJWT(routers.SaveTweet))).Methods(http.MethodPost)
	router.HandleFunc("/get-tweets", middleware.CheckBD(middleware.ValidateJWT(routers.GetTweets))).Methods(http.MethodGet)
	router.HandleFunc("/delete-tweet", middleware.CheckBD(middleware.ValidateJWT(routers.DeleteTweet))).Methods(http.MethodDelete)
	router.HandleFunc("/upload-avatar", middleware.CheckBD(middleware.ValidateJWT(routers.UploadAvatar))).Methods(http.MethodPost)
	router.HandleFunc("/upload-banner", middleware.CheckBD(middleware.ValidateJWT(routers.UploadBanner))).Methods(http.MethodPost)
	router.HandleFunc("/get-avatar", middleware.CheckBD(routers.GetAvatar)).Methods(http.MethodGet)
	router.HandleFunc("/get-banner", middleware.CheckBD(routers.GetBanner)).Methods(http.MethodGet)
	router.HandleFunc("/save-relation", middleware.CheckBD(middleware.ValidateJWT(routers.SaveRelation))).Methods(http.MethodPost)
	router.HandleFunc("/delete-relation", middleware.CheckBD(middleware.ValidateJWT(routers.DeleteRelation))).Methods(http.MethodDelete)
	router.HandleFunc("/get-relation", middleware.CheckBD(middleware.ValidateJWT(routers.GetRelation))).Methods(http.MethodGet)
	router.HandleFunc("/get-users", middleware.CheckBD(middleware.ValidateJWT(routers.GetAllUsers))).Methods(http.MethodGet)
	router.HandleFunc("/get-followers-tweets", middleware.CheckBD(middleware.ValidateJWT(routers.ReadFollowersTweets))).Methods(http.MethodGet)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
