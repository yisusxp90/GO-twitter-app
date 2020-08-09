package bd

import (
	"context"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTweet(tweet models.Tweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("tweet")

	register := bson.M{
		"userid":        tweet.UserID,
		"message":       tweet.Message,
		"creation_date": tweet.CreationDate,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return string(""), true, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
