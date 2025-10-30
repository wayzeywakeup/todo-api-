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

func CreateTask(title string) (Task, error) {
	res, err := db.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", title, false)
	if err != nil {
		return Task{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Task{}, err
	}

	task := Task{
		ID: int(id),
		Title: title,
		Done: false,
	}

	return task, nil
}

func UpdateTask(id int, done bool) (Task, error) {
	res, err := db.Exec("UPDATE tasks SET done = ? WHERE id = ?", id, done)
	if err != nil {
		return Task{}, err
	}

	ressAffected, err := res.RowsAffected()
	if err != nil {
		return Task{}, err
	}

	if ressAffected == 0 {
		return Task{}, fmt.Errorf("task with id %d not found", id)
	}

	var t Task

	err = db.QueryRow("SELECT id, title, done FROM tasks WHERE id = ?", id).Scan(t.ID, t.Title, t.Done)
	if err != nil {
		return Task{}, err
	}
	
	return t, nil
}

func DeleteTask(id int) error {
	res, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return  err
	}

	resAffected , err := res.RowsAffected()
	if err != nil {
		return  err
	}

	if resAffected == 0 {
		return fmt.Errorf("tasl with id %d not found", id)
	}

	return nil
}