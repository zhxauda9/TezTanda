package dal

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepo struct {
	db             *mongo.Database
	carsCollection *mongo.Collection
}

func NewCarRepo(client *mongo.Database) *CarRepo {
	return &CarRepo{db: client, carsCollection: client.Collection("Cars")}
}

// Add inserts a new car into the Cars collection.
func (r *CarRepo) Add(car Car) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	car.ID = primitive.NewObjectID() // Generate a new ObjectID for the car.
	_, err := r.carsCollection.InsertOne(ctx, car)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return car.ID, nil
}

// GetCar retrieves a single car by its ID.
func (r *CarRepo) GetCar(id primitive.ObjectID) (Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var car Car
	err := r.carsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&car)
	if err != nil {
		return Car{}, err
	}

	return car, nil
}

// GetCars retrieves all cars from the Cars collection.
func (r *CarRepo) GetCars() ([]Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.carsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var cars []Car
	if err = cursor.All(ctx, &cars); err != nil {
		return nil, err
	}

	return cars, nil
}

// Update modifies an existing car by its ID.
func (r *CarRepo) Update(id primitive.ObjectID, updatedCar Car) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"make":     updatedCar.Make,
			"price":    updatedCar.Price,
			"color":    updatedCar.Color,
			"model_id": updatedCar.ModelID,
		},
	}

	result, err := r.carsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no car found with the given ID")
	}

	return nil
}

// Delete removes a car by its ID.
func (r *CarRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := r.carsCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no car found with the given ID")
	}

	return nil
}
