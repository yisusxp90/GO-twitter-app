package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* Function to check if relation exist into bd*/
func GetRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("relation")

	// userid from table tweet in bd
	condition := bson.M{
		"userid":         relation.UserID,
		"userrelationid": relation.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	fmt.Println(result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
