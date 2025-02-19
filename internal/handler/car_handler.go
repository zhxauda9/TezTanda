package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Temutjin2k/CarTrading/internal/dal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CarHandler struct {
	car_repo *dal.CarRepo
}

func NewCarHandler(car_repo *dal.CarRepo) *CarHandler {
	return &CarHandler{car_repo: car_repo}
}

// AddNewCar handles adding a new car.
func (h *CarHandler) AddNewCar(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to add new car")

	var car dal.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		log.Println("Invalid input to add car")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.car_repo.Add(car)
	if err != nil {
		log.Println("Failed to add car: " + err.Error())
		http.Error(w, "Failed to add car: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Car added successfully",
		"id":      id.Hex(),
	})

	carToPrint, _ := json.MarshalIndent(car, "", "    ")
	log.Println("New car added!\nCar:", string(carToPrint))
}

// GetCars handles retrieving all cars.
func (h *CarHandler) GetCars(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for getting cars")

	cars, err := h.car_repo.GetCars()
	if err != nil {
		log.Println("Failed to fetch cars: " + err.Error())
		http.Error(w, "Failed to fetch cars: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request for getting cars handled succesfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

// GetCar handles retrieving a single car by ID.
func (h *CarHandler) GetCar(w http.ResponseWriter, r *http.Request) {

	carId := r.PathValue("id")
	log.Println("Request to get car by id:", carId)
	id, err := primitive.ObjectIDFromHex(carId)
	if err != nil {
		log.Println("Invalid car ID")
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	car, err := h.car_repo.GetCar(id)
	if err != nil {
		log.Println("Failed to fetch car: " + err.Error())
		http.Error(w, "Failed to fetch car: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

// UpdateCar handles updating a car by ID.
func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	carId := r.PathValue("id")
	log.Println("Request to update car id:", carId)

	id, err := primitive.ObjectIDFromHex(carId)
	if err != nil {
		log.Println("Invalid id:", carId)
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	var updatedCar dal.Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil {
		log.Println("Invalid input")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.car_repo.Update(id, updatedCar)
	if err != nil {
		log.Println("Failed to update car: " + err.Error())
		http.Error(w, "Failed to update car: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Car updated successfully"})
}

// DeleteCar handles deleting a car by ID.
func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {
	carId := r.PathValue("id")
	log.Println("Request to delete car with id", carId)

	id, err := primitive.ObjectIDFromHex(carId)
	if err != nil {
		log.Println("Invalid car ID:", carId)
		http.Error(w, "Invalid car ID", http.StatusBadRequest)
		return
	}

	err = h.car_repo.Delete(id)
	if err != nil {
		log.Println("Failed to delete car: " + err.Error())
		http.Error(w, "Failed to delete car: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Car deleted successfully"})
}
