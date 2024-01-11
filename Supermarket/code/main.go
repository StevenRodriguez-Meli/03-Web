package main

import (
	"app/Supermarket/code/handlers"
	"app/Supermarket/code/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	if err := handlers.LoadProducts(); err != nil {
		log.Fatalf("Error loading products: %s", err)
	}

	router := routes.SetupRoutes()

	fmt.Println("Server listening on :8080...")
	http.ListenAndServe(":8080", router)

}
