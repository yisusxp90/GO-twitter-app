package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

var Email string
var userID string

/* function to extract data from token */
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	secret := []byte("yisusxp-secret")
	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		// Generate Error, not must contain neither uppercase or signs
		return claims, false, string(""), errors.New("token format invalid")
	}
	token = strings.TrimSpace(splitToken[1])
	// save result into second parameter "claims"
	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	// token valid
	if err == nil {
		_, find, _ := bd.VerifyUserExist(claims.Email)
		if find == true {
			Email = claims.Email
			userID = claims.ID.Hex()
		}
		return claims, find, userID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
