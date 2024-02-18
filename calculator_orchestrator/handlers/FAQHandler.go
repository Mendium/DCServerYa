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
	fmt.Fprintf(w, "Operations:\n\"*\" requires 5ms,\n\"+\" requires 7ms (replaced with 'p'),\n\"-\" requires 3ms,\n\"/\" requires 8ms.")
}
