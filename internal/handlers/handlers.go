package handlers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/msuny-c/kode-rest-api/internal/database"
	"github.com/msuny-c/kode-rest-api/internal/helper"
	"github.com/msuny-c/kode-rest-api/internal/models"
	"github.com/msuny-c/kode-rest-api/internal/yaspeller"
	log "github.com/sirupsen/logrus"
)

var speller = yaspeller.New()

func CreateNote(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	note := r.Form.Get("note")
    typos, err := speller.Text(note)
    if len(typos) != 0 {
    	helper.WriteError(w, models.ResponseError{Code: http.StatusBadRequest, Message: "Found spelling errors"})
        return
    }
	err = database.CreateNote(username, note)
	if err != nil {
		log.Errorf("Failed to create note for user %q: %v.", username, err)
		helper.WriteError(w, models.ResponseError{Code: http.StatusInternalServerError, Message: "Internal Server Error"})
	} else {
		log.Infof("Created note from user %q.", username)
		helper.WriteResponse(w, models.Reponse{User: username, Note: note, Notes: nil})
	}
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	notes, err := database.ListNotes(username)
	if err != nil {
		log.Errorf("Failed to get notes from user %q: %v.", username, err)
		helper.WriteError(w, models.ResponseError{Code: http.StatusInternalServerError, Message: "Internal Server Error"})
	} else {
		helper.WriteResponse(w, models.Reponse{User: username, Note: "", Notes: notes})
	}
}