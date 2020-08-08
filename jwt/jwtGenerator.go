package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yisusxp90/GO-twitter-app/models"
)

func GenerateJWT(user models.Usuario) (string, error) {
	secret := []byte("yisusxp-secret")
	payload := jwt.MapClaims{
		"email":            user.Email,
		"nombre":           user.Nombre,
		"fecha_nacimiento": user.FechaNacimiento,
		"biografia":        user.Biografia,
		"ubicacion":        user.Ubicacion,
		"sitio_web":        user.SitioWeb,
		"_id":              user.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
