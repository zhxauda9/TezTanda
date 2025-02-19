package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Temutjin2k/CarTrading/internal/dal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ManufacturerHandler struct {
	Manufacturer_repo *dal.ManufactureRepo
}

func NewManufactureHandler(Manufacturer_repo *dal.ManufactureRepo) *ManufacturerHandler {
	return &ManufacturerHandler{Manufacturer_repo: Manufacturer_repo}
}

// AddManufacture handles adding a new manufacturer.
func (h *ManufacturerHandler) AddManufacture(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to add new manufacturer")

	var manufacturer dal.Manufacturer
	if err := json.NewDecoder(r.Body).Decode(&manufacturer); err != nil {
		log.Println("Invalid input to add manufacturer")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.Manufacturer_repo.Add(manufacturer)
	if err != nil {
		log.Println("Failed to add manufacturer: " + err.Error())
		http.Error(w, "Failed to add manufacturer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Manufacturer added successfully",
		"id":      id.Hex(),
	})

	manufacturerToPrint, _ := json.MarshalIndent(manufacturer, "", "    ")
	log.Println("New manufacturer added!\nManufacturer:", string(manufacturerToPrint))
}

// GetManufacture handles retrieving a manufacturer by ID.
func (h *ManufacturerHandler) GetManufacture(w http.ResponseWriter, r *http.Request) {
	manufacturerId := r.PathValue("id")
	log.Println("Request to get manufacturer by id:", manufacturerId)
	id, err := primitive.ObjectIDFromHex(manufacturerId)
	if err != nil {
		log.Println("Invalid manufacturer ID")
		http.Error(w, "Invalid manufacturer ID", http.StatusBadRequest)
		return
	}

	manufacturer, err := h.Manufacturer_repo.Get(id)
	if err != nil {
		log.Println("Failed to fetch manufacturer: " + err.Error())
		http.Error(w, "Failed to fetch manufacturer: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manufacturer)
}

// GetAllManufactures handles retrieving all manufacturers.
func (h *ManufacturerHandler) GetAllManufactures(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for getting all manufacturers")

	manufacturers, err := h.Manufacturer_repo.GetAll()
	if err != nil {
		log.Println("Failed to fetch manufacturers: " + err.Error())
		http.Error(w, "Failed to fetch manufacturers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request for getting all manufacturers handled successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manufacturers)
}

// UpdateManufacture handles updating a manufacturer by ID.
func (h *ManufacturerHandler) UpdateManufacture(w http.ResponseWriter, r *http.Request) {
	manufacturerId := r.PathValue("id")
	log.Println("Request to update manufacturer id:", manufacturerId)

	id, err := primitive.ObjectIDFromHex(manufacturerId)
	if err != nil {
		log.Println("Invalid id:", manufacturerId)
		http.Error(w, "Invalid manufacturer ID", http.StatusBadRequest)
		return
	}

	var updatedManufacturer dal.Manufacturer
	if err := json.NewDecoder(r.Body).Decode(&updatedManufacturer); err != nil {
		log.Println("Invalid input")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.Manufacturer_repo.Update(id, updatedManufacturer)
	if err != nil {
		log.Println("Failed to update manufacturer: " + err.Error())
		http.Error(w, "Failed to update manufacturer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Manufacturer updated successfully"})
}

// DeleteManufacture handles deleting a manufacturer by ID.
func (h *ManufacturerHandler) DeleteManufacture(w http.ResponseWriter, r *http.Request) {
	manufacturerId := r.PathValue("id")
	log.Println("Request to delete manufacturer with id", manufacturerId)

	id, err := primitive.ObjectIDFromHex(manufacturerId)
	if err != nil {
		log.Println("Invalid manufacturer ID:", manufacturerId)
		http.Error(w, "Invalid manufacturer ID", http.StatusBadRequest)
		return
	}

	err = h.Manufacturer_repo.Delete(id)
	if err != nil {
		log.Println("Failed to delete manufacturer: " + err.Error())
		http.Error(w, "Failed to delete manufacturer: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Manufacturer deleted successfully"})
}
