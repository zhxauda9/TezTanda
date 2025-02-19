package dal

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManufactureRepo struct {
	db                    *mongo.Database
	manufactureCollection *mongo.Collection
}

func NewManufactureRepo(client *mongo.Database) *ManufactureRepo {
	return &ManufactureRepo{db: client, manufactureCollection: client.Collection("Manufacturer")}
}

// Add inserts a new manufacturer into the Manufacturer collection.
func (r *ManufactureRepo) Add(manufacturer Manufacturer) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert manufacturer into the collection
	result, err := r.manufactureCollection.InsertOne(ctx, manufacturer)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Return the inserted ID
	return result.InsertedID.(primitive.ObjectID), nil
}

// GetManufacturer retrieves a single manufacturer by its ID.
func (r *ManufactureRepo) Get(id primitive.ObjectID) (Manufacturer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var manufacturer Manufacturer
	filter := bson.M{"_id": id}

	err := r.manufactureCollection.FindOne(ctx, filter).Decode(&manufacturer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Manufacturer{}, errors.New("manufacturer not found")
		}
		return Manufacturer{}, err
	}

	return manufacturer, nil
}

// GetManufacturers retrieves all manufacturers from the Manufacturer collection.
func (r *ManufactureRepo) GetAll() ([]Manufacturer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.manufactureCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var manufacturers []Manufacturer
	if err := cursor.All(ctx, &manufacturers); err != nil {
		return nil, err
	}

	return manufacturers, nil
}

// Update modifies an existing manufacturer by its ID.
func (r *ManufactureRepo) Update(id primitive.ObjectID, updatedManufacturer Manufacturer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":            updatedManufacturer.Name,
			"country":         updatedManufacturer.Country,
			"foundation_year": updatedManufacturer.FoundationYear,
			"website":         updatedManufacturer.Website,
		},
	}

	result, err := r.manufactureCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no manufacturer found with the given ID")
	}

	return nil
}

// Delete removes a manufacturer by its ID.
func (r *ManufactureRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := r.manufactureCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no manufacturer found with the given ID")
	}

	return nil
}
