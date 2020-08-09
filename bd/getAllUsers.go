package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* function to get either all users or only users we're following, this can be applyied with the param typeSearch, if receive follow will get only relationals users */
func GetAllUsers(ID string, page int64, search string, typeSearch string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("users")

	var results []*models.Usuario
	/* condition to search into BD _id:ObjectId("5f2eb1c75b0bd8235c6b930d") */
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	// (?i) search regardless neither uppercase or lowecase
	condition := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, condition, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	var find, include bool
	for cursor.Next(context.TODO()) {
		var user models.Usuario
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var relation models.Relation
		relation.UserID = ID
		relation.UserRelationID = user.ID.Hex()
		include = false
		find, err = GetRelation(relation)
		// get users i'm hot following
		if typeSearch == "new" && find == false {
			include = true
		}
		if typeSearch == "follow" && find == true {
			include = true
		}
		// if i'm following myself
		if relation.UserRelationID == ID {
			include = false
		}
		if include == true {
			user.Password = ""
			user.Biografia = ""
			user.SitioWeb = ""
			user.Ubicacion = ""
			user.Banner = ""
			user.Email = ""
			results = append(results, &user)
		}
	}
	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
