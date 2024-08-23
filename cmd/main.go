package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/msuny-c/kode-rest-api/internal/database"
	"github.com/msuny-c/kode-rest-api/internal/handlers"
	"github.com/msuny-c/kode-rest-api/internal/middleware"
	log "github.com/sirupsen/logrus"
)

func main() {
	database.ConnectDB()
	
	router := mux.NewRouter()
	
	router.Handle("/notes", middleware.Notes(http.HandlerFunc(handlers.CreateNote))).Methods("POST")
	router.HandleFunc("/notes", handlers.ListNotes).Methods("GET")
	
	amw := middleware.Authentication{Tokens: map[string]string{
		"wn29ANSMq39": "Alex",
		"k01NSmp2WQA": "Jane",
		"14mAf93Aq1f": "Kate",
	}}
	router.Use(amw.Middleware)
	
	log.Fatal(http.ListenAndServe(":8080", router))
}
