package dal

import (
	"context"
	"errors"
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

	existingProduct, _ := r.GetProduct(product.ID)
	if existingProduct.ID != primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("product already exists")
	}
	if product.Name == "" || product.Description == "" || product.Stock == 0 || product.Price == 0 || product.Category == "" {
		return primitive.NilObjectID, errors.New("please fill name, description, stock,price and category of product properly!")
	}
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

func (r *ProductRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := r.productCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

func (r *ProductRepo) Update(id primitive.ObjectID, updatedProduct Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	prevProduct, _ := r.GetProduct(updatedProduct.ID)

	if updatedProduct.Name == "" || updatedProduct.Description == "" || updatedProduct.Stock == 0 || updatedProduct.Price == 0 || updatedProduct.Category == "" {
		return errors.New("please fill name, description, stock,price and category of product properly!")
	}

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":        updatedProduct.Name,
			"description": updatedProduct.Description,
			"category":    updatedProduct.Category,
			"price":       updatedProduct.Price,
			"stock":       updatedProduct.Stock,
			"image":       updatedProduct.Image,
			"created_at":  prevProduct.CreatedAt,
			"updated_at":  time.Now(),
		},
	}

	result, err := r.productCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no product found with the given ID")
	}
	return nil
}
