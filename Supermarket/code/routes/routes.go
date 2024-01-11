package routes

import (
	"app/Supermarket/code/handlers"
	"net/http"

	"github.com/go-chi/chi"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", handlers.PingHandler)
	r.Get("/products", handlers.ProductsHandler)
	r.Get("/products/{id}", handlers.ProductByIDHandler)
	r.Get("/products/search", handlers.SearchProductsHandler)

	r.Post("/products", handlers.AddProductHandler)
	return r
}
