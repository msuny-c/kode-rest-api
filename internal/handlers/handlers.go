package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	"github.com/msuny-c/kode-rest-api/internal/database"
	log "github.com/sirupsen/logrus"
	
)

type Reponse struct {
	Notes []string `json:"notes"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	note := r.Form.Get("note")
	err := database.DB.CreateNote(username, note)
	if err != nil {
		log.Errorf("Failed to create note for user %q: %v.", username, err)
	}
	log.Infof("Created note from user %q.", username)
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	notes, err := database.DB.ListNotes(username)
	if err != nil {
		log.Errorf("Failed to get notes from user %q: %v.", username, err)
	}
	json, err := json.Marshal(Reponse{Notes: notes})
	if err != nil {
		log.Errorf("Failed to marshal notes to json: %v.", err)
	}
	w.Write(json)
}