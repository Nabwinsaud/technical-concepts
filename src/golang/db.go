package main

import (
	"database/sql"
	"fmt"
	"log"
)

// var db *sql.DB

func InitializeDB(host, port, username, password, database string) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, err
}
