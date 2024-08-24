package helper

import (
	"encoding/json"
	"net/http"
	"github.com/msuny-c/kode-rest-api/internal/models"
)

func WriteResponse(w http.ResponseWriter, response models.Response) {
	body, err := json.MarshalIndent(response, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(response.Code)
		w.Write(body)
	}
}