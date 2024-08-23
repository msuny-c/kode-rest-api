package database 

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type DBinstance struct {
	db *sql.DB
}

var DB DBinstance

func (storage *DBinstance) CreateNote(username string, note string) error {
	_, err := storage.db.Exec("INSERT INTO notes (note, username) VALUES ($1, $2)", note, username)
	if err != nil {
		log.Warnf("Couldn't add note to the database from user %q: %v", username, err)
		return err
	}
	return nil
}
func (storage *DBinstance) ListNotes(username string) ([]string, error) {
	var notes []string
	rows, err := storage.db.Query("SELECT note FROM notes WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var note string
		rows.Scan(&note)
		notes = append(notes, note)
	}
	return notes, nil
}

func ConnectDB() {
	dsn := fmt.Sprintf(
			"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS notes (id SERIAL PRIMARY KEY, note TEXT NOT NULL, username TEXT NOT NULL)")
	if err != nil {
		log.Fatalf("Failed to create required tables: %v", err)
	}
	DB = DBinstance{db: db}
}
