package handler

import (
	"TezTanda/internal/dal"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productRepo *dal.ProductRepo
}

func NewProductHandler(productRepo *dal.ProductRepo) *ProductHandler {
	return &ProductHandler{productRepo: productRepo}
}

func (h *ProductHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to add new product")

	err := r.ParseMultipartForm(10 << 20) // Ограничение: 10MB
	if err != nil {
		log.Println("Error parsing form data: ", err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	category := r.FormValue("category")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	file, header, err := r.FormFile("image")
	if err != nil {
		log.Println("Failed to read image: ", err)
		http.Error(w, "Failed to read image", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println("Failed to read image data: ", err)
		http.Error(w, "Failed to read image data", http.StatusInternalServerError)
		return
	}

	fileName := fmt.Sprintf("%s-%s", name, header.Filename)
	filePath, err := UploadToTripleS(fileBytes, fileName)
	if err != nil {
		log.Println("Failed to upload image: ", err)
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	product := dal.Product{
		Name:        name,
		Description: description,
		Category:    category,
		Price:       price,
		Stock:       stock,
		Image:       filePath,
	}

	id, err := h.productRepo.Add(product)
	if err != nil {
		log.Println("Failed to add product: " + err.Error())
		http.Error(w, "Failed to add product", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Product added successfully",
		"id":      id.Hex(),
	})
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for getting products")

	products, err := h.productRepo.GetProducts()
	if err != nil {
		log.Println("Failed to fetch products: " + err.Error())
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	log.Println("Request for getting products handled succesfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	log.Println("Request to get product by id:", productID)
	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Println("Invalid product ID")
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.GetProduct(id)
	if err != nil {
		log.Println("Failed to fetch product: " + err.Error())
		http.Error(w, "Failed to fetch product: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	log.Println("Request to update product id:", productID)

	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Println("Invalid id:", productID)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	existingProduct, err := h.productRepo.GetProduct(id)
	if err != nil {
		log.Println("Product not found: ", err)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // Ограничение: 10MB
	if err != nil {
		log.Println("Error parsing form data: ", err)
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	if name := r.FormValue("name"); name != "" {
		existingProduct.Name = name
	}
	if description := r.FormValue("description"); description != "" {
		existingProduct.Description = description
	}
	if category := r.FormValue("category"); category != "" {
		existingProduct.Category = category
	}
	if price := r.FormValue("price"); price != "" {
		existingProduct.Price, _ = strconv.ParseFloat(price, 64)
	}
	if stock := r.FormValue("stock"); stock != "" {
		existingProduct.Stock, _ = strconv.Atoi(stock)
	}

	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()

		if existingProduct.Image != "" {
			err = DeleteFromTripleS(existingProduct.Image)
			if err != nil {
				log.Println("Failed to delete old image: ", err)
			}
		}
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			log.Println("Failed to read image data: ", err)
			http.Error(w, "Failed to read image data", http.StatusInternalServerError)
			return
		}

		fileName := fmt.Sprintf("%s-%s", existingProduct.Name, header.Filename)
		filePath, err := UploadToTripleS(fileBytes, fileName)
		if err != nil {
			log.Println("Failed to update image: ", err)
			http.Error(w, "Failed to upload image", http.StatusInternalServerError)
			return
		}
		existingProduct.Image = filePath
	}

	err = h.productRepo.Update(id, existingProduct)
	if err != nil {
		log.Println("Failed to update product: " + err.Error())
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	log.Println("Request to delete product with id", productID)

	id, err := primitive.ObjectIDFromHex(productID)
	if err != nil {
		log.Println("Invalid product ID:", productID)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err = h.productRepo.Delete(id)
	if err != nil {
		log.Println("Failed to delete product: " + err.Error())
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}
