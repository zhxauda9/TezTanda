package dal

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartRepo struct {
	collection *mongo.Collection
}

func NewCartRepo(db *mongo.Database) *CartRepo {
	return &CartRepo{
		collection: db.Collection("cart"), // Коллекция "cart" в MongoDB
	}
}

func (r *CartRepo) AddProductToCart(productID string, quantity int) error {
	ctx := context.TODO()

	// Проверяем, есть ли уже товар в корзине
	filter := bson.M{"product_id": productID}
	var existingCartItem struct {
		ProductID string `bson:"product_id"`
		Quantity  int    `bson:"quantity"`
	}

	err := r.collection.FindOne(ctx, filter).Decode(&existingCartItem)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// Если товара нет в корзине — добавляем его
			_, err = r.collection.InsertOne(ctx, bson.M{
				"product_id": productID,
				"quantity":   quantity,
			})
			return err
		}
		return err
	}

	// Если товар уже есть — увеличиваем количество
	update := bson.M{"$set": bson.M{"quantity": existingCartItem.Quantity + quantity}}
	_, err = r.collection.UpdateOne(ctx, filter, update)
	return err
}
