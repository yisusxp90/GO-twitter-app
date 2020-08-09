package bd

import (
	"context"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadFollowersTweets(ID string, page int) ([]models.GetFollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("relation")

	skip := (page - 1) * 20
	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	// join table relation with tweet
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",          //this work like join between tables, tweet is the table
			"localField":   "userrelationid", // field to join the table
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"creation_date": -1}}) // sort desc
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	// with aggregate cursor is already done !! we don't have to iterate it
	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.GetFollowersTweets
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true

}
