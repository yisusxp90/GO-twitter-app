package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extention = strings.Split(handler.Filename, ".")[1]
	var fileName string = "uploads/banners/" + UserID + "." + extention
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error updating banner"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying banner"+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.Usuario
	var status bool
	user.Banner = UserID + "." + extention
	status, err = bd.UpdateProfile(user, UserID)
	if err != nil || status == false {
		http.Error(w, "Error saving banner into bd"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
