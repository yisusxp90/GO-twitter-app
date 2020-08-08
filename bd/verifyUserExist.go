package bd

import (
	"context"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* function to verify if user has already been registered*/
func VerifyUserExist(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel() // cancel when search into bd has finished
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("users")
	condition := bson.M{"email": email}
	var result models.Usuario
	// save user into &result
	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
