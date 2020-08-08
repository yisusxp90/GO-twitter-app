package main

/* locals packages*/
import (
	"log"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/handlers"
)

func main() {
	if bd.CheckConnection() == false {
		log.Fatal("no connection established to the bd")
	}
	handlers.Handlers()
}
