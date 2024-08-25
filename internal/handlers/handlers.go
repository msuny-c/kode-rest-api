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
var spellerCodes = map[int]string{
	1: "ERROR_UNKNOWN_WORD",
	2: "ERROR_REPEAT_WORD",
	3: "ERROR_CAPITALIZATION",
	4: "ERROR_TOO_MANY_ERRORS",
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
    username := context.Get(r, "username").(string)
    note := context.Get(r, "note").(string)
    typos, err := speller.Text(note)
    if len(typos) != 0 {
    	handleTypos(w, typos)
        return
    }
<<<<<<< HEAD
	err = database.CreateNote(username, note)
	if err != nil {
		log.Errorf("Failed to create note for user %q: %v.", username, err)
	} else {
		log.Infof("Created note from user %q.", username)
		helper.WriteResponse(w, http.StatusOK, models.ResponseNote{Note: note})
	}
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	username := context.Get(r, "username").(string)
	notes, err := database.ListNotes(username)
	if err != nil {
		log.Errorf("Failed to get notes from user %q: %v.", username, err)
	} else {
		helper.WriteResponse(w, http.StatusOK, models.ResponseNotes{Notes: notes})
	}
=======
    err = database.CreateNote(username, note)
    if err != nil {
	log.Errorf("Failed to create note for user %q: %v.", username, err)
    } else {
	log.Infof("Created note from user %q.", username)
	helper.WriteResponse(w, models.Response{Code: http.StatusOK, User: username, Note: note})
    }
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
    username := context.Get(r, "username").(string)
    notes, err := database.ListNotes(username)
    if err != nil {
	log.Errorf("Failed to get notes from user %q: %v.", username, err)
    } else {
	helper.WriteResponse(w, models.Response{Code: http.StatusOK, User: username, Notes: &notes})
    }
>>>>>>> b0ea1888b83db4dc7f3ff2f58475389e9baad36a
}
	
func handleTypos(w http.ResponseWriter, typos []yaspeller.Response) {
<<<<<<< HEAD
	var errors models.Errors
=======
    var errors []models.Error
>>>>>>> b0ea1888b83db4dc7f3ff2f58475389e9baad36a
    for _, typo := range typos {
    	errors.Add(spellerCodes[typo.Code], "Spelling error in word: " + typo.Word)
    }
<<<<<<< HEAD
    helper.WriteResponse(w, http.StatusBadRequest, models.ResponseError{Errors: errors})
}
=======
    helper.WriteResponse(w, models.Response{Code: http.StatusBadRequest, Errors: errors})
}
>>>>>>> b0ea1888b83db4dc7f3ff2f58475389e9baad36a
