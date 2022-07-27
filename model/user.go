package model

import (
	"log"
)

func CreateUser(user []string) {
	db := connect()
	defer db.Close()
	_, err := db.Query("INSERT into users(first_name,last_name,email,hash) VALUES ($1,$2,$3,$4)", user[0], user[1], user[2], user[3])
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
}

func SelectHash(email string) string {
	db := connect()
	defer db.Close()
	row, err := db.Query("Select hash from users where email =$1", email)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
	var hash string
	row.Next()
	row.Scan(&hash)
	return hash
}
