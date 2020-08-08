package middleware

import (
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
)

/* Middleware to check the connection bd status*/
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == false {
			http.Error(w, "BD lost connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
