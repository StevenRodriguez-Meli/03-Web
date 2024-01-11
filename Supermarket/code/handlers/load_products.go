package handlers

import (
	"app/Supermarket/code/models"
	"encoding/json"
	"os"
)

var productIDCounter int
var products []models.Product

func LoadProducts() error {
	fileData, err := os.ReadFile("products.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(fileData, &products); err != nil {
		return err
	}

	for _, product := range products {
		id := product.ID
		if id > productIDCounter {
			productIDCounter = id
		}
	}

	return nil
}
