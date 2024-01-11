package handlers

import (
	"app/Supermarket/code/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	sendJSONResponse(w, products)
}

func ProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	println("idParam: ", idParam)
	id, err := strconv.Atoi(idParam)
	println("id: ", id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product := findProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	sendJSONResponse(w, product)
}

func SearchProductsHandler(w http.ResponseWriter, r *http.Request) {
	priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)
	if err != nil {
		http.Error(w, "Invalid priceGt parameter", http.StatusBadRequest)
		return
	}

	filteredProducts := filterProductsByPrice(priceGt)
	sendJSONResponse(w, filteredProducts)
}

func AddProductHandler(w http.ResponseWriter, r *http.Request) {
	var productRequest models.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productRequest); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if err := validateProductRequest(&productRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	nextID := getNextProductID()

	newProduct := models.Product{
		ID:          nextID,
		Name:        productRequest.Name,
		Quantity:    productRequest.Quantity,
		CodeValue:   productRequest.CodeValue,
		IsPublished: productRequest.IsPublished,
		Expiration:  productRequest.Expiration,
		Price:       productRequest.Price,
	}

	products = append(products, newProduct)

	sendJSONResponse(w, newProduct)
}
