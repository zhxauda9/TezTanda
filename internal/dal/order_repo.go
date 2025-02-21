package dal

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// OrderRepo — репозиторий для работы с заказами
type OrderRepo struct {
	collection *mongo.Collection
}

// NewOrderRepo создает новый экземпляр OrderRepo
func NewOrderRepo(db *mongo.Database) *OrderRepo {
	return &OrderRepo{
		collection: db.Collection("orders"),
	}
}

// CreateOrder добавляет новый заказ в базу данных
func (r *OrderRepo) CreateOrder(ctx context.Context, order *Order) (primitive.ObjectID, error) {
	order.ID = primitive.NewObjectID()
	order.CreatedAt = time.Now()
	order.Status = "pending" // Новый заказ всегда в статусе "pending"

	fmt.Println("Saving order:", order) // Лог перед вставкой

	result, err := r.collection.InsertOne(ctx, order)
	if err != nil {
		fmt.Println("Error inserting order:", err)
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		fmt.Println("Inserted ID is not ObjectID:", result.InsertedID)
		return primitive.NilObjectID, fmt.Errorf("failed to convert InsertedID")
	}

	fmt.Println("Order inserted successfully, ID:", insertedID)
	return insertedID, nil
}

// GetOrders возвращает все заказы
func (r *OrderRepo) GetOrders(ctx context.Context) ([]Order, error) {
	var orders []Order
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("Error fetching orders:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order Order
		if err := cursor.Decode(&order); err != nil {
			fmt.Println("Error decoding order:", err)
			continue
		}
		orders = append(orders, order)
	}

	fmt.Println("Fetched orders:", orders)
	return orders, nil
}
