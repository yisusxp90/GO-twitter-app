package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

/* function to get Banner, this api won't have response*/
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")
	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Page must be greater than 0", http.StatusBadRequest)
		return
	}
	pag := int64(pageTemp)
	result, status := bd.GetAllUsers(UserID, pag, search, typeUser)
	if status == false {
		http.Error(w, "Error getting users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
