package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

/* function to save tweet into BD */
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.TweetDecode
	err := json.NewDecoder(r.Body).Decode(&message)
	register := models.Tweet{
		UserID:       UserID,
		Message:      message.Message,
		CreationDate: time.Now(),
	}
	_, status, err := bd.SaveTweet(register)
	if err != nil {
		http.Error(w, "Error saving tweet "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "Fail saving user ", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(register)
}
