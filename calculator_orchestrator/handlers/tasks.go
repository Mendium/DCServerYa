package handlers

import (
	"DistributedComputingServer/calculator_agent"
	"DistributedComputingServer/calculator_orchestrator/models"
	"DistributedComputingServer/database"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTasksHandler(w, r)
	case http.MethodPost:
		createTaskHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		http.Error(w, "Failed to open database connection", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var status, expression string
	query := "SELECT status, expression FROM tasks WHERE id = ?"
	err = db.QueryRow(query, id).Scan(&status, &expression)
	if err != nil {
		http.Error(w, "Failed to query task status", http.StatusInternalServerError)
		return
	}
	if status == "ready" {
		fmt.Fprintf(w, "Your expression has been calculated. Result: %s", expression)
	} else {
		fmt.Fprintf(w, "Your expression is still being calculated.")
	}
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	expression := r.Form.Get("expression")

	expression = strings.ReplaceAll(expression, "p", "+")

	newTask := models.Task{
		Answer: expression,
		Status: "pending",
	}

	id, err := database.AddTaskToDatabase(newTask)
	if err != nil {
		http.Error(w, "Failed to add task to database", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Task added to database with ID: %d", id)
	err = calculator_agent.AgentDo(expression, id)
	if err != nil {
		return
	}
}
