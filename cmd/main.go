package main

import (
	"TezTanda/internal/defaultData"
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"TezTanda/internal/dal"
	"TezTanda/internal/handler"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	SetupLogger()
	log.Println("Trying to load .env file...")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")

	log.Println("Trying to connect mongoDB...")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	log.Println("Successfuly connected to MongoDB!")

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	fsUploads := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fsUploads))

	db := client.Database("TezTanda")
	defaultData.Fill(db)

	user_repo := dal.NewUserRepo(db)
	user_handler := handler.NewUserHandler(user_repo)
	auth_handler := handler.NewAuthHandler(user_repo)

	mux.HandleFunc("/", ServePage)
	mux.HandleFunc("POST /users", user_handler.AddNewUser)
	mux.HandleFunc("POST /login", auth_handler.Login)
	mux.HandleFunc("POST /logout", auth_handler.Logout)
	mux.HandleFunc("GET /users", user_handler.GetUsers)
	mux.HandleFunc("GET /users/{id}", user_handler.GetUser)
	mux.HandleFunc("GET /profile", auth_handler.GetProfile)
	mux.HandleFunc("PUT /users/{id}", user_handler.UpdateUser)
	mux.HandleFunc("DELETE /users/{id}", user_handler.DeleteUser)

	product_repo := dal.NewProductRepo(db)
	product_handler := handler.NewProductHandler(product_repo)

	mux.HandleFunc("POST /products", product_handler.AddProduct)
	mux.HandleFunc("GET /products", product_handler.GetProducts)
	mux.HandleFunc("GET /products/{id}", product_handler.GetProduct)
	mux.HandleFunc("PUT /products/{id}", product_handler.UpdateProduct)
	mux.HandleFunc("DELETE /products/{id}", product_handler.DeleteProduct)

	log.Println("Server started on address: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func SetupLogger() {
	logFileName := "server.log"

	_ = os.Remove(logFileName)

	logFile, err := os.Create(logFileName)
	if err != nil {
		log.Fatalf("Error to create log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Logging is set up! The recording goes to the file and to the terminal.")
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "web/home.html")
		return
	}

	filePath := "web" + r.URL.Path
	_, err := os.Stat(filePath)

	if err == nil {
		http.ServeFile(w, r, filePath)
	} else {
		http.NotFound(w, r)
	}
}