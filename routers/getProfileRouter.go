package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

/* Get profile values*/
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// get id from url
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id can not be null", http.StatusBadRequest)
		return
	}
	profile, err := bd.GetProfile(ID)
	if err != nil {
		http.Error(w, "User not found"+err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
