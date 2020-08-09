package models

/* decode tweet from body */
type TweetDecode struct {
	Message string `bson:"message" json:"message"`
}
