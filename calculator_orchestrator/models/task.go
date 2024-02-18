package models

type Task struct {
	ID     int    `json:"id"`
	Answer string `json:"answer"`
	Status string `json:"status"`
}
