package dal

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarModelRepo struct {
	db                 *mongo.Database
	carModelCollection *mongo.Collection
}

func NewCarModelRepo(client *mongo.Database) *CarModelRepo {
	return &CarModelRepo{db: client, carModelCollection: client.Collection("Model")}
}

// Add inserts a new model into the Model collection.
func (r *CarModelRepo) Add(carModel Model) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	carModel.ID = primitive.NewObjectID() // Generate a new ObjectID for the model.
	_, err := r.carModelCollection.InsertOne(ctx, carModel)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return carModel.ID, nil
}

// GetCar retrieves a single model by its ID.
func (r *CarModelRepo) Get(id primitive.ObjectID) (Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var carModel Model
	err := r.carModelCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&carModel)
	if err != nil {
		return Model{}, err
	}

	return carModel, nil
}

// GetCars retrieves all models from the Cars collection.
func (r *CarModelRepo) GetAll() ([]Model, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.carModelCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var models []Model
	if err = cursor.All(ctx, &models); err != nil {
		return nil, err
	}

	return models, nil
}

// Update modifies an existing model by its ID.

func (r *CarModelRepo) Update(id primitive.ObjectID, updatedModel Model) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":            updatedModel.Name,
			"country":         updatedModel.Country,
			"year":            updatedModel.Year,
			"description":     updatedModel.Description,
			"manufacturer_id": updatedModel.ManufacturerID,
		},
	}

	result, err := r.carModelCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no model found with the given ID")
	}

	return nil
}

// Delete removes a model by its ID.
func (r *CarModelRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := r.carModelCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no carModel found with the given ID")
	}

	return nil
}
