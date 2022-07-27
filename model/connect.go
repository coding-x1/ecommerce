package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//pgAdmin 4 used for management
//=============================
func connect() *sql.DB {
	connStr := "postgresql://postgres:secret@localhost:5432/entry?sslmode=disable"

	// Connect to database
	db, _ := sql.Open("postgres", connStr)
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
