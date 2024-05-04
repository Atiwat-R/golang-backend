package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body") // If there's no requrst body, return Error
	}
	return json.NewDecoder(r.Body).Decode(payload) // Decode JSON into payload
}

func WriteJSON(w http.ResponseWriter, status_code int, v any) error {
	w.Header().Add("Content-Type", "application/json") 
	w.WriteHeader(status_code)
	return json.NewEncoder(w).Encode(v)
}

// Write an Error to User
func WriteError(w http.ResponseWriter, status_code int, err error) {
	// Map [string:string], to store the JSON we'll return
	responseJSON  := map[string]string {
		"error" : err.Error(),
	}
	WriteJSON(w, status_code, responseJSON)
}
