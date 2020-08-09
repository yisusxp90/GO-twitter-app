package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* struct to get all tweets */
type GetFollowersTweets struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	Tweet          struct {
		Message   string             `bson:"message" json:"message,omitempty"`
		TweetDate time.Time          `bson:"tweetdate" json:"tweetDate,omitempty"`
		ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	}
}
