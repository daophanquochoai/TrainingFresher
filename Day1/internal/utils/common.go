package utils

import (
	"encoding/json"
	"go-rest-api/internal/model"
	"net/http"
)

// write request
func WriteJson(w http.ResponseWriter, statusCode int, response model.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
