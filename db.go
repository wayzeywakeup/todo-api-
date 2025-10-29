package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func initDB() {
	database, err := sql.Open("sqlite", "./task.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	statement, err := database.Prepare(
		`CREATE TABLE IF NOT EXISTS task(
			id INTEGER PRIMARY KEY,
			title TEXT,
			done BOOLEAN
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized successfully")
}