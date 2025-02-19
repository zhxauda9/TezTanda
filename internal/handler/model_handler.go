package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Temutjin2k/CarTrading/internal/dal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ModelHandler struct {
	model_repo *dal.CarModelRepo
}

func NewModelHandler(model_repo *dal.CarModelRepo) *ModelHandler {
	return &ModelHandler{model_repo: model_repo}
}

// AddModel handles adding a new model.
func (h *ModelHandler) AddModel(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to add new model")

	var model dal.Model
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		log.Println("Invalid input to add model")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.model_repo.Add(model)
	if err != nil {
		log.Println("Failed to add model: " + err.Error())
		http.Error(w, "Failed to add model: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Model added successfully",
		"id":      id.Hex(),
	})

	modelToPrint, _ := json.MarshalIndent(model, "", "    ")
	log.Println("New model added!\nModel:", string(modelToPrint))
}

// GetModel handles retrieving a single model by ID.
func (h *ModelHandler) GetModel(w http.ResponseWriter, r *http.Request) {
	modelId := r.PathValue("id")
	log.Println("Request to get model by id:", modelId)

	id, err := primitive.ObjectIDFromHex(modelId)
	if err != nil {
		log.Println("Invalid model ID")
		http.Error(w, "Invalid model ID", http.StatusBadRequest)
		return
	}

	model, err := h.model_repo.Get(id)
	if err != nil {
		log.Println("Failed to fetch model: " + err.Error())
		http.Error(w, "Failed to fetch model: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model)
}

// GetModels handles retrieving all models.
func (h *ModelHandler) GetModels(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to get all models")

	models, err := h.model_repo.GetAll()
	if err != nil {
		log.Println("Failed to fetch models: " + err.Error())
		http.Error(w, "Failed to fetch models: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request to get all models handled successfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}

// UpdateModel handles updating a model by ID.
func (h *ModelHandler) UpdateModel(w http.ResponseWriter, r *http.Request) {
	modelId := r.PathValue("id")
	log.Println("Request to update model id:", modelId)

	id, err := primitive.ObjectIDFromHex(modelId)
	if err != nil {
		log.Println("Invalid model ID:", modelId)
		http.Error(w, "Invalid model ID", http.StatusBadRequest)
		return
	}

	var updatedModel dal.Model
	if err := json.NewDecoder(r.Body).Decode(&updatedModel); err != nil {
		log.Println("Invalid input")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.model_repo.Update(id, updatedModel)
	if err != nil {
		log.Println("Failed to update model: " + err.Error())
		http.Error(w, "Failed to update model: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Model updated successfully"})
}

// DeleteModel handles deleting a model by ID.
func (h *ModelHandler) DeleteModel(w http.ResponseWriter, r *http.Request) {
	modelId := r.PathValue("id")
	log.Println("Request to delete model with id:", modelId)

	id, err := primitive.ObjectIDFromHex(modelId)
	if err != nil {
		log.Println("Invalid model ID:", modelId)
		http.Error(w, "Invalid model ID", http.StatusBadRequest)
		return
	}

	err = h.model_repo.Delete(id)
	if err != nil {
		log.Println("Failed to delete model: " + err.Error())
		http.Error(w, "Failed to delete model: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Model deleted successfully"})
}
