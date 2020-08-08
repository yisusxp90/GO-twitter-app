package bd

import (
	"context"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveUser(user models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("users")
	user.Password, _ = EncryptPassword(user.Password)
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
