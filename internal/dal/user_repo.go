package dal

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db             *mongo.Database
	userCollection *mongo.Collection
}

func NewUserRepo(client *mongo.Database) *UserRepo {
	return &UserRepo{db: client, userCollection: client.Collection("Users")}
}

// Add inserts a new user into the User collection.
func (r *UserRepo) Add(user User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	existingUser, _ := r.GetUser(user.ID)
	if existingUser.ID != primitive.NilObjectID {
		return primitive.NilObjectID, errors.New("user already exists")
	}

	if user.Name == "" || user.Email == "" {
		return primitive.NilObjectID, errors.New("name and email cannot be empty")
	}
	user.ID = primitive.NewObjectID() // Generate a new ObjectID for the user.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.Password = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = r.userCollection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return user.ID, nil
}

// Update modifies an existing user by its ID.
func (r *UserRepo) Update(id primitive.ObjectID, updatedUser User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	prevUse, _ := r.GetUser(updatedUser.ID)

	if updatedUser.Name == "" || updatedUser.Email == "" {
		return errors.New("name and email cannot be empty")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"name":       updatedUser.Name,
			"surname":    updatedUser.Surname,
			"email":      updatedUser.Email,
			"password":   hashedPassword,
			"role":       updatedUser.Role,
			"created_at": prevUse.CreatedAt,
			"updated_at": time.Now(),
		},
	}

	result, err := r.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}

// GetUser retrieves a single user by its ID.
func (r *UserRepo) GetUser(id primitive.ObjectID) (User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user User
	err := r.userCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// Get Users retrieves all users from the User collection.
func (r *UserRepo) GetUsers() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	result, err := r.userCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no user found with the given ID")
	}

	return nil
}
