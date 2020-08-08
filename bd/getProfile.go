package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("users")

	var profile models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)
	/* condition to search into BD _id:ObjectId("5f2eb1c75b0bd8235c6b930d") */
	condition := bson.M{
		"_id": objID,
	}
	// get and save data into profile
	err := col.FindOne(ctx, condition).Decode(&profile)
	// not show password
	profile.Password = ""
	if err != nil {
		fmt.Println("Register not found" + err.Error())
		return profile, err
	}

	return profile, nil
}
