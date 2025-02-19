package dal

import "go.mongodb.org/mongo-driver/bson/primitive"

// the Cars collection
type Car struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Make    string             `bson:"make" json:"make"`
	Price   float64            `bson:"price" json:"price"`
	Color   string             `bson:"color" json:"color"`
	ModelID primitive.ObjectID `bson:"model_id" json:"model_id"`
}

// the Models collection
type Model struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Country        string             `bson:"country" json:"country"`
	Year           int                `bson:"year" json:"year"`
	Description    string             `bson:"description" json:"description"`
	ManufacturerID primitive.ObjectID `bson:"manufacturer_id" json:"manufacturer_id"`
}

// the Manufacturers collection
type Manufacturer struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name" json:"name"`
	Country        string             `bson:"country" json:"country"`
	FoundationYear int                `bson:"foundation_year" json:"foundation_year"`
	Website        string             `bson:"website" json:"website"`
}
