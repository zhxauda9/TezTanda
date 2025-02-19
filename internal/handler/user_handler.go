package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"TezTanda/internal/dal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	user_repo *dal.UserRepo
}

func NewUserHandler(user_repo *dal.UserRepo) *UserHandler {
	return &UserHandler{user_repo: user_repo}
}

// AddNewUser handles adding a new user.
func (h *UserHandler) AddNewUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Request to add new user")

	var user dal.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Invalid input to add user")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	id, err := h.user_repo.Add(user)
	if err != nil {
		log.Println("Failed to add user: " + err.Error())
		http.Error(w, "Failed to add user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User added successfully",
		"id":      id.Hex(),
	})

	userToPrint, _ := json.MarshalIndent(user, "", "    ")
	log.Println("New user added!\nUser:", string(userToPrint))
}

// GetUsers handles retrieving all users.
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Request for getting users")

	users, err := h.user_repo.GetUsers()
	if err != nil {
		log.Println("Failed to fetch users: " + err.Error())
		http.Error(w, "Failed to fetch users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Request for getting user handled succesfully")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser handles retrieving a single user by ID.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	userID := r.PathValue("id")
	log.Println("Request to get user by id:", userID)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Invalid user ID")
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.user_repo.GetUser(id)
	if err != nil {
		log.Println("Failed to fetch user: " + err.Error())
		http.Error(w, "Failed to fetch user: "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser handles updating a user by ID.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	usedId := r.PathValue("id")
	log.Println("Request to update user id:", usedId)

	id, err := primitive.ObjectIDFromHex(usedId)
	if err != nil {
		log.Println("Invalid id:", usedId)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updatedUser dal.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		log.Println("Invalid input")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.user_repo.Update(id, updatedUser)
	if err != nil {
		log.Println("Failed to update user: " + err.Error())
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

// DeleteUser handles deleting a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("id")
	log.Println("Request to delete user with id", userID)

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println("Invalid user ID:", userID)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.user_repo.Delete(id)
	if err != nil {
		log.Println("Failed to delete user: " + err.Error())
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
