package bd

import (
	"github.com/yisusxp90/GO-twitter-app/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.Usuario, bool) {
	user, find, _ := VerifyUserExist(email)
	if find == false {
		return user, false
	}
	passordBytes := []byte(password)
	passordBD := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passordBD, passordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
