package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func ParseJson(r *http.Request, payload interface{}) error {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Println("Error closing body")
		}
	}()
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return err
	}
	return nil
}

func WriteJsonError(w http.ResponseWriter, status int, err error) {
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
