package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "page is required", http.StatusBadRequest)
		return
	}
	// Atoi: convert char to int
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page must be major than 0", http.StatusBadRequest)
		return
	}
	pag := int64(page)
	result, valid := bd.GetTweets(ID, pag)
	if valid == false {
		http.Error(w, "Error getting tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
