package helper

import (
	"encoding/json"
	"net/http"
	"github.com/msuny-c/kode-rest-api/internal/models"
)

func WriteError(w http.ResponseWriter, error models.ResponseError) {
	body, err := json.MarshalIndent(error, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(error.Code)
		w.Write(body)
	}
}

func WriteResponse(w http.ResponseWriter, response models.Reponse) {
	body, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(body)
	}
}