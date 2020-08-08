package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/jwt"
	"github.com/yisusxp90/GO-twitter-app/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var userBody models.Usuario
	/* important - r.Body is a only read Stream object that's to say you can read this only once, save body into &user */
	err := json.NewDecoder(r.Body).Decode(&userBody)
	if err != nil {
		http.Error(w, "User & Password invalid. "+err.Error(), 400)
		return
	}
	if len(userBody.Email) == 0 {
		http.Error(w, "Field email is required", 400)
		return
	}
	userBD, userExist := bd.Login(userBody.Email, userBody.Password)
	if userExist == false {
		http.Error(w, "User & Password invalid. ", 400)
		return
	}
	// generate JWT
	jwtKey, err := jwt.GenerateJWT(userBD)
	if err != nil {
		http.Error(w, "Error generating jwt. "+err.Error(), 400)
		return
	}
	resp := models.LoginResponse{
		Token: jwtKey,
	}
	// set response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
