package models

/* Relation into BD */
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`                 // user id
	UserRelationID string `bson:"userrelationid" json:"userRelationId"` // user id from user i'm following
}
