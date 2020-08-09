package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

/* function to save relation into BD, receive by parameters userId i'm gonna follow */
func SaveRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is Required", http.StatusBadRequest)
		return
	}
	var relation models.Relation
	relation.UserID = UserID
	relation.UserRelationID = ID

	status, err := bd.SaveRelation(relation)
	if err != nil || status == false {
		http.Error(w, "it happened a mistake saving relation", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(relation)
}
