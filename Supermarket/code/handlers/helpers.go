package handlers

import (
	"app/Supermarket/code/models"
	"fmt"
	"time"
)

func filterProductsByPrice(priceGt float64) []models.Product {
	var filtered []models.Product
	for _, p := range products {
		if p.Price > priceGt {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func findProductByID(id int) *models.Product {
	for _, p := range products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

func validateProductRequest(product *models.ProductRequest) error {

	if product.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	if product.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	if product.CodeValue == "" {
		return fmt.Errorf("CodeValue cannot be empty")
	}

	if isCodeValueExists(product.CodeValue) {
		return fmt.Errorf("CodeValue already exists")
	}

	_, err := time.Parse("01/02/2006", product.Expiration)
	if err != nil {
		return fmt.Errorf("invalid ExpiryDate format. Use MM/DD/YYYY")
	}

	if product.Price < 0 {
		return fmt.Errorf("price cannot be negative")
	}
	return nil

}
func getNextProductID() int {
	productIDCounter++
	return productIDCounter
}

func isCodeValueExists(codeValue string) bool {
	for _, p := range products {
		if p.CodeValue == codeValue {
			return true
		}
	}
	return false
}
