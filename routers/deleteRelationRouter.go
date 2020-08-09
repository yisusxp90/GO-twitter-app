package routers

import (
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

/* function to delete relation from bd, receive by params userId i'm gonna leave to follow*/
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id parameter is Required", http.StatusBadRequest)
		return
	}
	var relation models.Relation
	relation.UserID = UserID
	relation.UserRelationID = ID
	status, err := bd.DeleteRelation(relation)
	if err != nil || status == false {
		http.Error(w, "it happend a mistake deleting the relation "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
