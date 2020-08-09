package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

/* function to get Banner, this api won't have response*/
func GetBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is Required", http.StatusBadRequest)
		return
	}
	profile, err := bd.GetProfile(ID)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}
	openFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "image not found", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error coping image", http.StatusBadRequest)
	}
}
