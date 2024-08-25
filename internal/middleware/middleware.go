package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/context"
	"github.com/msuny-c/kode-rest-api/internal/helper"
	"github.com/msuny-c/kode-rest-api/internal/models"
	log "github.com/sirupsen/logrus"
)

type Authentication struct {
	Tokens map[string]string
}

func (amw *Authentication) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if user, found := amw.Tokens[token]; found {
			log.Infof("Authenticated user: %q.", user)
			context.Set(r, "username", amw.Tokens[token])
			next.ServeHTTP(w, r)
		} else {
			log.Warnf("Unsuccessful authentication attempt from %s.", r.RemoteAddr)
			errors := []models.Error{{Code: "UNAUTHORIZED", Message: "X-Session-Token is invalid"}}
			helper.WriteResponse(w, http.StatusBadRequest, models.ResponseError{Errors: errors})
		}
	})
}

func Notes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request models.Request
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Errorf("Failed to parse request's body: %v", err)
		}
		json.Unmarshal(body, &request)
		if request.Note != "" {
			context.Set(r, "note", request.Note)
			next.ServeHTTP(w, r)
		} else {
			log.Infof("Required key 'note' was not provided from %s.", r.RemoteAddr)
			errors := []models.Error{{Code: "BAD REQUEST", Message: "Required key 'note' was not provided"}}
			helper.WriteResponse(w, http.StatusBadRequest, models.ResponseError{Errors: errors})
		}
	})
}