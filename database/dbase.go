package database

import (
	"DistributedComputingServer/calculator_orchestrator/models"
	"database/sql"
	"fmt"
	"log"
)

func AddTaskToDatabase(task models.Task) (int, error) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	query := "INSERT INTO tasks (expression, status) VALUES ($1, $2) RETURNING id"

	var id int
	err = db.QueryRow(query, task.Answer, task.Status).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to add task to database: %v", err)
	}

	return id, nil
}
