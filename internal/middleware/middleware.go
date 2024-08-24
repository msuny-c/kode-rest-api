package middleware

import (
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
		w.Header().Set("Content-Type", "application/json")
		token := r.Header.Get("X-Session-Token")
		if user, found := amw.Tokens[token]; found {
			log.Infof("Authenticated user: %q.", user)
			context.Set(r, "username", amw.Tokens[token])
			next.ServeHTTP(w, r)
		} else {
			log.Warnf("Unsuccessful authentication attempt from %s.", r.RemoteAddr)
			errors := []models.Error{{Code: "UNAUTHORIZED", Message: "X-Session-Token is invalid"}}
			helper.WriteError(w, models.ResponseError{Code: http.StatusUnauthorized, Errors: errors})
		}
	})
}

func Notes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Errorf("Error occured while parsing body: %v.", err)
		}
		if r.Form.Has("note") {
			next.ServeHTTP(w, r)
		} else {
			log.Infof("Required key 'note' was not provided from %s.", r.RemoteAddr)
			errors := []models.Error{{Code: "BAD REQUEST", Message: "Required key 'note' was not provided"}}
			helper.WriteError(w, models.ResponseError{Code: http.StatusBadRequest, Errors: errors})
		}
	})
}