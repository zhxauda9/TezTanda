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
		// Dairy 21
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
		{ID: primitive.NewObjectID(), Name: "Greek yougurt", Description: "Greek yogurt Food Master natural 8,4%, 130 g, Kazakhstan", Category: "Dairy", Price: 355, Stock: 200, Image: "uploads/greek.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Egg", Description: "Egg Kazger Couscous, 30 pieces, Kazakhstan", Category: "Dairy", Price: 2220, Stock: 200, Image: "uploads/eggs.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Yogurt", Description: "Food Master drinking yogurt with wild berry flavor 2%, 900 ml, Kazakhstan", Category: "Dairy", Price: 745, Stock: 200, Image: "uploads/foodmaster.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Irimshik", Description: "Irimshik rafaello sweet Tandauly by galmart, weight, Kazakhstan", Category: "Dairy", Price: 2115, Stock: 200, Image: "uploads/irimshik.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Turkish ayran", Description: "Turkish Natige ayran 1,4% 250g, Kazakhstan", Category: "Dairy", Price: 250, Stock: 200, Image: "uploads/turkish.jfif", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Miracle", Description: "Miracle glazed Coconut cheese 23%, 40 g, Russia", Category: "Dairy", Price: 380, Stock: 200, Image: "uploads/miracle.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "President", Description: "Sour cream President 15%, 200 g, Kazakhstan", Category: "Dairy", Price: 490, Stock: 200, Image: "uploads/president.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Suluguni", Description: "Natige white suluguni cheese 45%, 350 g, Kazakhstan", Category: "Dairy", Price: 2975, Stock: 200, Image: "uploads/suluguni.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Mozarella", Description: "Galbani Mozzarella Santa Lucia cheese 48%, 125 g, Russia", Category: "Dairy", Price: 1080, Stock: 200, Image: "uploads/mozarella.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Epica", Description: "Yogurt Epica cherry-sweet cherry 4,8%, 130 ml, Russia", Category: "Dairy", Price: 449, Stock: 200, Image: "uploads/epica.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Condensed milk", Description: "Condensed milk Dep 8.5%, 600 g, Kazakhstan", Category: "Dairy", Price: 380, Stock: 200, Image: "uploads/condensed.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Fruits 11
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
		{ID: primitive.NewObjectID(), Name: "Avocado", Description: "Fresh avocado, 1kg", Category: "Fruits", Price: 1080, Stock: 200, Image: "uploads/avocado.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Vegetables 14
		{ID: primitive.NewObjectID(), Name: "Tomatoes", Description: "Fresh red tomatoes, 1kg", Category: "Vegetables", Price: 990, Stock: 80, Image: "uploads/tomatoes.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Potatoes", Description: "Young potatoes, 1kg", Category: "Vegetables", Price: 389, Stock: 100, Image: "uploads/potatoes.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Carrots", Description: "Organic carrots, 1kg", Category: "Vegetables", Price: 180, Stock: 70, Image: "uploads/carrots.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Onions", Description: "Fresh onions, 1kg", Category: "Vegetables", Price: 120, Stock: 80, Image: "uploads/onions.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Garlics", Description: "Fresh garlics, 1kg", Category: "Vegetables", Price: 1190, Stock: 80, Image: "uploads/garlic.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cucumbers", Description: "Fresh cucumbers, 1kg", Category: "Vegetables", Price: 1290, Stock: 80, Image: "uploads/cucumbers.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Ginger", Description: "Fresh ginger, 1kg", Category: "Vegetables", Price: 1890, Stock: 80, Image: "uploads/ginger.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pepper", Description: "Fresh pepper, 1kg", Category: "Vegetables", Price: 1490, Stock: 80, Image: "uploads/pepper.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Zuccini", Description: "Zucchini is fresh, weight", Category: "Vegetables", Price: 780, Stock: 200, Image: "uploads/zuccini.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Radish", Description: "Fresh radish, bunch, Kazakhstan", Category: "Vegetables", Price: 550, Stock: 200, Image: "uploads/radish.jpeg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pumpkin", Description: "Pumpkin, weight, Kazakhstan", Category: "Vegetables", Price: 337, Stock: 200, Image: "uploads/pumpkin.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Beetroot", Description: "Beetroot, weight, Kazakhstan", Category: "Vegetables", Price: 117, Stock: 200, Image: "uploads/beetroot.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Beijin Cabbage", Description: "Peking cabbage mini, weight, Kazakhstan", Category: "Vegetables", Price: 998, Stock: 200, Image: "uploads/peking.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Cauliflower", Description: "Cauliflower, weight", Category: "Vegetables", Price: 1320, Stock: 200, Image: "uploads/caulifllower.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Eggplant", Description: "Eggplant NAC-AGRO, weight, Kazakhstan", Category: "Vegetables", Price: 840, Stock: 200, Image: "uploads/eggplant.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Daikon", Description: "Daikon radish, weight", Category: "Vegetables", Price: 324, Stock: 200, Image: "uploads/daikon.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Mushrooms", Description: "Oyster mushrooms, weight, Kazakhstan", Category: "Vegetables", Price: 1260, Stock: 200, Image: "uploads/mushrooms.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Leek", Description: "Fresh leek, 1000 g, Kazakhstan", Category: "Vegetables", Price: 2400, Stock: 200, Image: "uploads/leek.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Meat 16
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
		{ID: primitive.NewObjectID(), Name: "Salmon", Description: "Fresh salmon fillet, 500g", Category: "Meat", Price: 5300, Stock: 25, Image: "uploads/salmon.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Beef Bijan", Description: "Beef Bijan delicately sliced, in stock, 200 g, Kazakhstan", Category: "Meat", Price: 2005, Stock: 200, Image: "uploads/bijan.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Et Bayram", Description: "Boiled sausage Et Bayram Halal Muslim mini, 600 g, Kazakhstan", Category: "Meat", Price: 1450, Stock: 200, Image: "uploads/bayram.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Narlen", Description: "Boiled milk sausages, 460 g, Kazakhstan", Category: "Meat", Price: 1070, Stock: 200, Image: "uploads/narlen.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Shrimps", Description: "Agama royal shrimp No. 4 for pasta, peeled, 300 g, Russia", Category: "Meat", Price: 6410, Stock: 200, Image: "uploads/shrimp.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Drinks 16
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
		{ID: primitive.NewObjectID(), Name: "Sprite", Description: "Sprite, 500ml", Category: "Drinks", Price: 500, Stock: 200, Image: "uploads/sprite.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pepsi", Description: "Pepsi, 500ml", Category: "Drinks", Price: 500, Stock: 200, Image: "uploads/pepsi.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Dr.Pepper", Description: "Dr.Pepper, 500ml", Category: "Drinks", Price: 500, Stock: 200, Image: "uploads/drpepper.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Fanta", Description: "Fanta, 500ml", Category: "Drinks", Price: 500, Stock: 200, Image: "uploads/fanta.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Natakhtari Saperavi", Description: "Natakhtari Saperavi, 500ml", Category: "Drinks", Price: 670, Stock: 200, Image: "uploads/saperavi.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Rich Juice", Description: "Rich Juice, 500ml", Category: "Drinks", Price: 380, Stock: 200, Image: "uploads/rich.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},

		// Bakery 20
		{ID: primitive.NewObjectID(), Name: "Cottage cheese bun", Description: "Galmart cottage cheese bun, weight, Kazakhstan", Category: "Bakery", Price: 380, Stock: 200, Image: "uploads/cottageCheese.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Sour cream bun", Description: "Galmart Sour cream bun, 90 g, pcs, Kazakhstan", Category: "Bakery", Price: 190, Stock: 200, Image: "uploads/creamBun.jpeg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Mini bun with poppy seeds", Description: "Mini bun with poppy seeds galmart, in, Kazakhstan", Category: "Bakery", Price: 100, Stock: 200, Image: "uploads/miniBun.jfif", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Apple bun", Description: "Apple bun 1/80, piece, Kazakhstan", Category: "Bakery", Price: 180, Stock: 200, Image: "uploads/appleBun.jfif", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Hamburger bun", Description: "Hamburger bun 60g, Kazakhstan", Category: "Bakery", Price: 100, Stock: 200, Image: "uploads/hamburgerBun.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Rose bun", Description: "Galmart Rose bun, pcs, Kazakhstan", Category: "Bakery", Price: 190, Stock: 200, Image: "uploads/roseBun.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pastry pretzel", Description: "Galmart pastry pretzel, pcs, Kazakhstan", Category: "Bakery", Price: 280, Stock: 200, Image: "uploads/pastry.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Meat patty", Description: "Shortbread meat patty galmart", Category: "Bakery", Price: 380, Stock: 200, Image: "uploads/meatPatty.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Samsa dastarkhan", Description: "Samsa dastarkhan", Category: "Bakery", Price: 400, Stock: 200, Image: "uploads/samsaDas.jfif", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Croissant with cinnamon", Description: "Croissant with cinnamon galmart, pieces, Kazakhstan", Category: "Bakery", Price: 250, Stock: 200, Image: "uploads/croissantCinnamon.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Puff pastry", Description: "Puff pastry with cottage cheese galmart, in, Kazakhstan", Category: "Bakery", Price: 260, Stock: 200, Image: "uploads/puffPastry.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Milk Puff Pastry", Description: "Puff pastry with condensed milk galmart, in, Kazakhstan", Category: "Bakery", Price: 280, Stock: 200, Image: "uploads/milkPuffPastry.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Baursaks", Description: "Galmart baursaks, weight, Kazakhstan", Category: "Bakery", Price: 1120, Stock: 200, Image: "uploads/bauyrsak.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Sausages in dough", Description: "Sausages baked in galmart dough, 90 g, Kazakhstan", Category: "Bakery", Price: 410, Stock: 200, Image: "uploads/sausages.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chocolate croissant", Description: "Croissant with chocolate galmart, in, Kazakhstan", Category: "Bakery", Price: 380, Stock: 200, Image: "uploads/croissantChocolate.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chebureks", Description: "Chebureks with meat galmart, 100 g, pcs, Kazakhstan", Category: "Bakery", Price: 400, Stock: 200, Image: "uploads/chebureks.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pirozhki", Description: "Fried pies with onion and egg galmart, 60 g, Kazakhstan", Category: "Bakery", Price: 180, Stock: 200, Image: "uploads/pirozhki1.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Pirozhki", Description: "Fried liver pies galmart, in, Kazakhstan", Category: "Bakery", Price: 190, Stock: 200, Image: "uploads/pirozhki2.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Chuck-chuck", Description: "Chuck-Chuck Tatar galmart, weight, Kazakhstan", Category: "Bakery", Price: 400, Stock: 200, Image: "uploads/chuckchuck.jpeg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Sandwich", Description: "Croissant sandwich with sausage 1/150 galmart, pcs, Kazakhstan", Category: "Bakery", Price: 1000, Stock: 200, Image: "uploads/sandwich.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: primitive.NewObjectID(), Name: "Loaf", Description: "Galmart Embassy loaf, pcs, Kazakhstan", Category: "Bakery", Price: 380, Stock: 200, Image: "uploads/loaf.jpg", CreatedAt: time.Now(), UpdatedAt: time.Now()},
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
