package auth

import (
	"encoding/json"
	"net/http"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Login successful",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
