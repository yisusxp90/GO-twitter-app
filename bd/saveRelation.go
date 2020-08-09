package bd

import (
	"context"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
)

func SaveRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("relation")
	_, err := col.InsertOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}
