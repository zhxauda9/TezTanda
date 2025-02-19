package dal

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProductRepo struct {
	db                *mongo.Database
	productCollection *mongo.Collection
}

func NewProductRepo(client *mongo.Database) *ProductRepo {
	return &ProductRepo{db: client, productCollection: client.Collection("Products")}
}

func (r *ProductRepo) Add(product Product) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	product.ID = primitive.NewObjectID()
	_, err := r.productCollection.InsertOne(ctx, product)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return product.ID, nil
}

func (r *ProductRepo) GetProduct(id primitive.ObjectID) (Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product Product
	err := r.productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *ProductRepo) GetProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.productCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []Product
	if err = cursor.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}
