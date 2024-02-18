package main

import (
	"DistributedComputingServer/calculator_orchestrator/handlers"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/tasks", handlers.TasksHandler)
	http.HandleFunc("/operations", handlers.OperataionsHandler)

	fmt.Println("Server listening on port 8080...")
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, expression TEXT, status TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
