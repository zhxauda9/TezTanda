package handler

import (
	"TezTanda/internal/dal"
	"encoding/json"
	"net/http"
)

type ProductHandler struct {
	productRepo *dal.ProductRepo
}

func NewProductHandler(productRepo *dal.ProductRepo) *ProductHandler {
	return &ProductHandler{productRepo: productRepo}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product dal.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.productRepo.Add(product)
	if err != nil {
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product added successfully",
		"id":      id.Hex(),
	})
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.GetProducts()
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
