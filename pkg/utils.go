package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MessageError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func WriteError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := MessageError{Status: status, Message: message}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, `"message": "failed to encode error response"`, http.StatusInternalServerError)
	}
}

func ParseJSON(payload any, r *http.Request) error {
	if r.Body == nil {
		return fmt.Errorf("error: invalid json format")
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return fmt.Errorf("error: unable to encode json")
	}

	return nil
}

func WriteJSON(w http.ResponseWriter, status int, output any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(output); err != nil {
		return fmt.Errorf("error: encoding json")
	}

	return nil
}
