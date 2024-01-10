package greetings

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Greeting struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GreetingsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data Greeting

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if data.FirstName == "" || data.LastName == "" {
		http.Error(w, "Missing first name or last name", http.StatusBadRequest)
		return
	}

	response := GenerateResponse(data)
	SendResponse(w, response)
}

func GenerateResponse(data Greeting) string {
	return fmt.Sprintf("Hello %s %s", data.FirstName, data.LastName)
}

func SendResponse(w http.ResponseWriter, response string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(response))
}
