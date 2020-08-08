package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

/* function to create into bd the user register */
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	/* important - r.Body is a only read Stream object that's to say you can read this only once, save body into &user */
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error into receive data "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "field Email is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Password must have at least 6 characters", 400)
		return
	}

	/* verify email isn't registered*/
	_, find, _ := bd.VerifyUserExist(user.Email)
	if find == true {
		http.Error(w, "user is already registered", 400)
		return
	}
	_, status, err := bd.SaveUser(user)
	if err != nil {
		http.Error(w, "Error saving user "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Fail saving user ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
