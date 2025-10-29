package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDB() {
	var err error

	db, err = sql.Open("sqlite", "./task.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS tasks(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT,
			done BOOLEAN
		)`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database initialized successfully")
}

func GetTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, title, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var t Task

		err := rows.Scan(&t.ID, &t.Title, &t.Done)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}