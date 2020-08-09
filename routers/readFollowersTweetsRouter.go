package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

/* function to save relation into BD, receive by parameters userId i'm gonna follow */
func ReadFollowersTweets(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if len(page) < 1 {
		http.Error(w, "page parameter is Required", http.StatusBadRequest)
		return
	}
	pg, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "page must be greater than 0", http.StatusBadRequest)
		return
	}

	resp, valid := bd.ReadFollowersTweets(UserID, pg)
	if valid == false {
		http.Error(w, "it happened a mistake reading tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
