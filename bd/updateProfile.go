package bd

import (
	"context"
	"time"

	"github.com/yisusxp90/GO-twitter-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateProfile(user models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/* cancel each timeout to not create spaces into memorie */
	defer cancel()
	db := MongoCN.Database("Twitter-app")
	col := db.Collection("users")
	// function make to create slices or maps key: string && values type: interfaces
	profile := make(map[string]interface{})
	if len(user.Nombre) > 0 {
		profile["nombre"] = user.Nombre
	}
	if len(user.Apellidos) > 0 {
		profile["apellidos"] = user.Apellidos
	}
	profile["fechaNacimiento"] = user.FechaNacimiento
	if len(user.Avatar) > 0 {
		profile["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		profile["banner"] = user.Banner
	}
	if len(user.Biografia) > 0 {
		profile["biografia"] = user.Biografia
	}
	if len(user.Ubicacion) > 0 {
		profile["ubicacion"] = user.Ubicacion
	}
	if len(user.SitioWeb) > 0 {
		profile["sitioWeb"] = user.SitioWeb
	}
	updateString := bson.M{
		"$set": profile,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
