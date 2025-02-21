package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"TezTanda/internal/dal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderHandler — обработчик заказов
type OrderHandler struct {
	orderRepo *dal.OrderRepo
}

// NewOrderHandler создает новый OrderHandler
func NewOrderHandler(orderRepo *dal.OrderRepo) *OrderHandler {
	return &OrderHandler{orderRepo: orderRepo}
}

// AddOrder — обработчик создания заказа
func (h *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	// Считываем body в строку для логирования
	body, _ := io.ReadAll(r.Body)
	log.Println("📩 Полученный JSON:", string(body))

	// Декодируем JSON-запрос
	var reqBody struct {
		UserID string `json:"user_id"`
		Items  []struct {
			ProductID string  `json:"product_id"`
			Quantity  int     `json:"quantity"`
			Price     float64 `json:"price"`
		} `json:"items"`
	}
	if err := json.Unmarshal(body, &reqBody); err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	// Проверяем, что `user_id` передан
	if reqBody.UserID == "" {
		log.Println("❌ Ошибка: user_id отсутствует в JSON")
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	userID, err := primitive.ObjectIDFromHex(reqBody.UserID)
	if err != nil {
		log.Println("❌ Ошибка конвертации user_id:", err)
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	// Формируем заказ
	order := dal.Order{
		ID:        primitive.NewObjectID(),
		UserID:    userID,
		Items:     []dal.OrderItem{},
		TotalCost: 0,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// Проверяем товары и рассчитываем сумму заказа
	for i, item := range reqBody.Items {
		if item.ProductID == "" {
			log.Println("❌ Ошибка: product_id отсутствует в позиции", i)
			http.Error(w, "product_id is required for each item", http.StatusBadRequest)
			return
		}

		productID, err := primitive.ObjectIDFromHex(item.ProductID)
		if err != nil {
			log.Println("❌ Ошибка конвертации product_id:", err)
			http.Error(w, "Invalid product ID format", http.StatusBadRequest)
			return
		}

		order.Items = append(order.Items, dal.OrderItem{
			ProductID: productID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		})
		order.TotalCost += float64(item.Quantity) * item.Price
	}

	// Добавляем заказ в базу
	insertedID, err := h.orderRepo.CreateOrder(context.Background(), &order)
	if err != nil {
		log.Println("❌ Ошибка создания заказа в БД:", err)
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Order created successfully",
		"orderID": insertedID.Hex(),
	})
}
