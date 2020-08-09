package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

/* function to update into bd the user register */
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	/* important - r.Body is a only read Stream object that's to say you can read this only once, save body into &user */
	err := json.NewDecoder(r.Body).Decode(&user)

	var status bool
	status, err = bd.UpdateProfile(user, UserID)

	if err != nil {
		http.Error(w, "Error updating user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Fail updating user ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
