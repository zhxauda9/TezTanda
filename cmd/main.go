package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Temutjin2k/CarTrading/internal/dal"
	"github.com/Temutjin2k/CarTrading/internal/handler"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	log.Println("Trying to load .env file...")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	log.Println("Trying to connect mongoDB Atlas...")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Println("Successfuly connected to MongoDB Atlas!")

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	car_collection := client.Database("CarSharing")
	car_repo := dal.NewCarRepo(car_collection)
	car_handler := handler.NewCarHandler(car_repo)

	mux.HandleFunc("/", ServePage)
	mux.HandleFunc("POST /cars", car_handler.AddNewCar)
	mux.HandleFunc("GET /cars", car_handler.GetCars)
	mux.HandleFunc("GET /cars/{id}", car_handler.GetCar)
	mux.HandleFunc("PUT /cars/{id}", car_handler.UpdateCar)
	mux.HandleFunc("DELETE /cars/{id}", car_handler.DeleteCar)

	model_repo := dal.NewCarModelRepo(car_collection)
	model_handler := handler.NewModelHandler(model_repo)

	mux.HandleFunc("POST /models", model_handler.AddModel)
	mux.HandleFunc("GET /models", model_handler.GetModels)
	mux.HandleFunc("GET /models/{id}", model_handler.GetModel)
	mux.HandleFunc("PUT /models/{id}", model_handler.UpdateModel)
	mux.HandleFunc("DELETE /models/{id}", model_handler.DeleteModel)

	manu_repo := dal.NewManufactureRepo(car_collection)
	manu_handler := handler.NewManufactureHandler(manu_repo)

	mux.HandleFunc("POST /manufacturers", manu_handler.AddManufacture)
	mux.HandleFunc("GET /manufacturers", manu_handler.GetAllManufactures)
	mux.HandleFunc("GET /manufacturers/{id}", manu_handler.GetManufacture)
	mux.HandleFunc("PUT /manufacturers/{id}", manu_handler.UpdateManufacture)
	mux.HandleFunc("DELETE /manufacturers/{id}", manu_handler.DeleteManufacture)

	log.Println("Server started on address: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/index.html")
}
