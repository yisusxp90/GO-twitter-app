package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN  connection object to BD  */
var MongoCN = connectBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://yisusxp90:yisusxp90@twitter-app.o9udn.mongodb.net/Twitter-app?retryWrites=true&w=majority")

/* connectBD function that allow to connect with the BD */
func connectBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("connection successfully")
	return client
}

/* CheckConnection pint to bd */
func CheckConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}
