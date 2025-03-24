package handlers

import (
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		HandleGetUser(w, r)
	case http.MethodPost:
		HandleNewUser(w, r)
	case http.MethodPut:
		HandleUpdateUser(w, r)
	case http.MethodDelete:
		HandleDeleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
