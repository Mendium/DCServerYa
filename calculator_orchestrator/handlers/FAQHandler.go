package handlers

import (
	"fmt"
	"net/http"
)

func OperataionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getOperationsHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getOperationsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Operations:\n\"*\" requires 10000ms,\n\"+\" requires 2000ms (replaced with 'p'),\n\"-\" requires 2500ms,\n\"/\" requires 12500ms.")
}
