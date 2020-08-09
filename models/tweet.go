package models

import "time"

/* Tweet into BD */
type Tweet struct {
	UserID       string    `bson:"userid" json:"userid,omitempty"`
	Message      string    `bson:"message" json:"message,omitempty"`
	CreationDate time.Time `bson:"creation_date" json:"creation_date,omitempty"`
}
