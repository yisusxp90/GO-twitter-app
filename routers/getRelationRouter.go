package routers

import (
	"encoding/json"
	"net/http"

	"github.com/yisusxp90/GO-twitter-app/bd"
	"github.com/yisusxp90/GO-twitter-app/models"
)

/* Get relation*/
func GetRelation(w http.ResponseWriter, r *http.Request) {
	// get id from url
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id can not be null", http.StatusBadRequest)
		return
	}
	var relation models.Relation
	relation.UserID = UserID
	relation.UserRelationID = ID

	var resp models.ResponseRelationresult

	status, err := bd.GetRelation(relation)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
