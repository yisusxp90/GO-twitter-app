package routers

import (
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is Required", http.StatusBadRequest)
		return
	}
	err := bd.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "it happend a mistake deleting the tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
