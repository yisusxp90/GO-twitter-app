package bd

import (
	"context"
	"log"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Function to get all tweets, parameters: userId and page*/
func GetTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("tweet")
	var result []*models.GetTweets
	// userid from table tweet in bd
	condition := bson.M{
		"userid": ID,
	}
	opts := options.Find()
	opts.SetLimit(20)
	opts.SetSort(bson.D{{Key: "creation_date", Value: -1}})
	opts.SetSkip((page - 1) * 20)
	cursor, err := col.Find(ctx, condition, opts)
	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}
	for cursor.Next(context.TODO()) {
		var register models.GetTweets
		err := cursor.Decode(&register)
		if err != nil {
			return result, false
		}
		result = append(result, &register)
	}
	return result, true
}
