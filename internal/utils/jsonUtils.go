package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var Validate = validator.New()

func ParseJson(r *http.Request, payload interface{}) error {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Println("Error closing body")
		}
	}()

	if r.Body == http.NoBody {
		return fmt.Errorf("request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJson(w http.ResponseWriter, status int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		return err
	}
	return nil
}

func WriteJsonError(w http.ResponseWriter, status int, err error) {
	log.Println("Error:", err)
	WriteJson(w, status, map[string]string{"error": err.Error()})
}
