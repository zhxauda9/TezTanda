package handler

import (
	"TezTanda/internal/dal"
	"encoding/json"
	"net/http"
)

type CartHandler struct {
	cartRepo *dal.CartRepo
}

func NewCartHandler(cartRepo *dal.CartRepo) *CartHandler {
	return &CartHandler{cartRepo: cartRepo}
}

func (h *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.cartRepo.AddProductToCart(request.ProductID, request.Quantity)
	if err != nil {
		http.Error(w, "Failed to add product to cart", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Product added to cart successfully"})
}
