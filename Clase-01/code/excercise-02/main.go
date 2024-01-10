package main

import (
	"app/Clase-01/code/excercise-02/greetings"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Post("/greetings", greetings.GreetingsHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
