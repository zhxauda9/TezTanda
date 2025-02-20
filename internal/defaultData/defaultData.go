package defaultData

import (
	"context"
	"log"
	"time"

	"TezTanda/internal/dal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Fill(client *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := client.Collection("Users")
	sana, _ := bcrypt.GenerateFromPassword([]byte("sana"), bcrypt.DefaultCost)
	aray, _ := bcrypt.GenerateFromPassword([]byte("aray"), bcrypt.DefaultCost)
	aida, _ := bcrypt.GenerateFromPassword([]byte("aida"), bcrypt.DefaultCost)

	users := []dal.User{
		{ID: primitive.NewObjectID(), Name: "Sana", Surname: "Bagym", Email: "sana@gmail.com", Password: string(sana), Role: "user", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Aray", Surname: "Bidanova", Email: "aray@gmail.com", Password: string(aray), Role: "admin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Aida", Surname: "Zhalgassova", Email: "aida@gmail.com", Password: string(aida), Role: "admin", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, user := range users {
		if err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Err(); err == mongo.ErrNoDocuments {
			_, err = userCollection.InsertOne(ctx, user)
			if err != nil {
				log.Println("Error inserting default user:", user.Name)
			}
		}
	}
	log.Println("User default data filled successfully!")

	productCollection := client.Collection("Products")

	products := []dal.Product{
		// Dairy Products 10
		{ID: primitive.NewObjectID(), Name: "Milk", Description: "Fresh cow milk, 1L", Category: "Dairy", Price: 590, Stock: 50, Image: "uploads/milk.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cheese", Description: "Hard cheese, 200g", Category: "Dairy", Price: 590, Stock: 40, Image: "uploads/cheese.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Yogurt", Description: "Strawberry yogurt, 150g", Category: "Dairy", Price: 390, Stock: 60, Image: "uploads/yogurt.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Butter", Description: "Salted butter, 250g", Category: "Dairy", Price: 780, Stock: 35, Image: "uploads/butter.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cream", Description: "Heavy cream, 500ml", Category: "Dairy", Price: 890, Stock: 30, Image: "uploads/cream.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Kefir", Description: "Tasty kefir, 500ml", Category: "Dairy", Price: 680, Stock: 30, Image: "uploads/kefir.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Sour Cream", Description: "Sour cream( smetana ), 400g", Category: "Dairy", Price: 890, Stock: 30, Image: "uploads/sourcream.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cocktail", Description: "Strawberry milk cocktail", Category: "Dairy", Price: 450, Stock: 30, Image: "uploads/cocktail.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cottage Cheese", Description: "Cottage cheese", Category: "Dairy", Price: 560, Stock: 30, Image: "uploads/cottage.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Qurt", Description: "Salty qurt", Category: "Dairy", Price: 390, Stock: 30, Image: "uploads/qurt.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Fruits 10
		{ID: primitive.NewObjectID(), Name: "Apples", Description: "Green apples, 1kg", Category: "Fruits", Price: 670, Stock: 75, Image: "uploads/apples.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Bananas", Description: "Ripe bananas, 1kg", Category: "Fruits", Price: 990, Stock: 90, Image: "uploads/bananas.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Oranges", Description: "Fresh oranges, 1kg", Category: "Fruits", Price: 990, Stock: 60, Image: "uploads/oranges.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Grapes", Description: "Sweet seedless grapes, 500g", Category: "Fruits", Price: 1290, Stock: 50, Image: "uploads/grapes.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Mandarins", Description: "Fresh mandarins, 1kg", Category: "Fruits", Price: 1190, Stock: 20, Image: "uploads/mandarin.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pears", Description: "Tasty pears, 1kg", Category: "Fruits", Price: 1290, Stock: 75, Image: "uploads/pears.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Promenade", Description: "Promenade 1kg", Category: "Fruits", Price: 1590, Stock: 75, Image: "uploads/promenade.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Limons", Description: "Tasty limons, 1kg", Category: "Fruits", Price: 1690, Stock: 75, Image: "uploads/limon.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Grapefruit", Description: "Grapefruit 1kg", Category: "Fruits", Price: 1890, Stock: 75, Image: "uploads/grapefruit.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Kiwi", Description: "Kiwi 1kg", Category: "Fruits", Price: 990, Stock: 75, Image: "uploads/kiwis.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Vegetables 10
		{ID: primitive.NewObjectID(), Name: "Tomatoes", Description: "Fresh red tomatoes, 1kg", Category: "Vegetables", Price: 990, Stock: 80, Image: "uploads/tomatoes.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Potatoes", Description: "Young potatoes, 1kg", Category: "Vegetables", Price: 389, Stock: 100, Image: "uploads/potatoes.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Carrots", Description: "Organic carrots, 1kg", Category: "Vegetables", Price: 180, Stock: 70, Image: "uploads/carrots.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Onions", Description: "Fresh onions, 1kg", Category: "Vegetables", Price: 120, Stock: 80, Image: "uploads/onions.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Garlics", Description: "Fresh garlics, 1kg", Category: "Vegetables", Price: 1190, Stock: 80, Image: "uploads/garlic.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cucumbers", Description: "Fresh cucumbers, 1kg", Category: "Vegetables", Price: 1290, Stock: 80, Image: "uploads/cucumbers.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Ginger", Description: "Fresh ginger, 1kg", Category: "Vegetables", Price: 1890, Stock: 80, Image: "uploads/ginger.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pepper", Description: "Fresh pepper, 1kg", Category: "Vegetables", Price: 1490, Stock: 80, Image: "uploads/pepper.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Meat & Fish 12
		{ID: primitive.NewObjectID(), Name: "Chicken thighs", Description: "Boneless skinless chicken thighs, 1kg", Category: "Meat", Price: 1860, Stock: 40, Image: "uploads/chicken_thighs.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chicken drumstick", Description: "Boneless skinless chicken drumstick, 1kg", Category: "Meat", Price: 2790, Stock: 40, Image: "uploads/chicken_drumstick.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chicken fillet", Description: "Boneless skinless chicken fillet, 1kg", Category: "Meat", Price: 2590, Stock: 40, Image: "uploads/chicken_fillet.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chicken wings", Description: "Boneless skinless chicken wings, 1kg", Category: "Meat", Price: 2620, Stock: 40, Image: "uploads/chicken_wings.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chicken legs", Description: "Boneless skinless chicken wings, 1kg", Category: "Meat", Price: 2370, Stock: 40, Image: "uploads/chicken_legs.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Minced chicken", Description: "Boneless skinless minced chicken, 1kg", Category: "Meat", Price: 2690, Stock: 40, Image: "uploads/chicken_minced.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Ground beef", Description: "Ground beef, 1kg", Category: "Meat", Price: 3050, Stock: 40, Image: "uploads/beef_grounded.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Horse kazy", Description: "Horse meat kazy, 1kg", Category: "Meat", Price: 5980, Stock: 40, Image: "uploads/kazy.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Beef meat", Description: "Beaf meat, 1kg", Category: "Meat", Price: 6500, Stock: 40, Image: "uploads/beef.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Lamb ribs", Description: "Lamb ribs, 1kg", Category: "Meat", Price: 5300, Stock: 40, Image: "uploads/lamb_ribs.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Lamb loin", Description: "Lamb loin, 1kg", Category: "Meat", Price: 4900, Stock: 40, Image: "uploads/lamb_loin.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Salmon", Description: "Fresh salmon fillet, 500g", Category: "Fish", Price: 5300, Stock: 25, Image: "uploads/salmon.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Drinks 10
		{ID: primitive.NewObjectID(), Name: "Water", Description: "Mineral water, 1L", Category: "Drinks", Price: 250, Stock: 200, Image: "uploads/water.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cola", Description: "Cola, 500ml", Category: "Drinks", Price: 500, Stock: 40, Image: "uploads/cola.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Fuse-tea Peach", Description: "Fuse-tea Peach, 500ml", Category: "Drinks", Price: 400, Stock: 70, Image: "uploads/fuse_peach.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Fuse-tea Berry", Description: "Fuse-tea Berry, 500ml", Category: "Drinks", Price: 400, Stock: 70, Image: "uploads/fuse_berry.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Fuse-tea Limon", Description: "Fuse-tea Limon, 500ml", Category: "Drinks", Price: 400, Stock: 70, Image: "uploads/fuse_limon.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Fuse-tea Dandelion", Description: "Fuse-tea Dandelion, 500ml", Category: "Drinks", Price: 400, Stock: 70, Image: "uploads/fuse_dand.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Fuse-tea Lime", Description: "Fuse-tea Lime, 500ml", Category: "Drinks", Price: 400, Stock: 70, Image: "uploads/fuse_lime.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Red Bull", Description: "Red Bull, 500ml", Category: "Drinks", Price: 800, Stock: 70, Image: "uploads/redbull.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Gorilla", Description: "Gorilla, 500ml", Category: "Drinks", Price: 500, Stock: 70, Image: "uploads/gorilla.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Apple Juice", Description: "Apple Juice, 500ml", Category: "Drinks", Price: 650, Stock: 70, Image: "uploads/applejuice.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	for _, product := range products {
		if err := productCollection.FindOne(ctx, bson.M{"name": product.Name}).Err(); err == mongo.ErrNoDocuments {
			_, err := productCollection.InsertOne(ctx, product)
			if err != nil {
				log.Println("Error inserting default product:", product.Name)
			}
		}
	}
	log.Println("Product default data filled successfully!")
}
