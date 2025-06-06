package helpers

import (
	"encoding/json"
	"net/http"
)

// RespondJSON writes a JSON response with given status code
func NetHttpJsonResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}