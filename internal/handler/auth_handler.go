package handler

import (
	"TezTanda/internal/dal"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

type AuthHandler struct {
	userRepo *dal.UserRepo
}

func NewAuthHandler(userRepo *dal.UserRepo) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, fmt.Sprintf("Invalid input: %v", err), http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, fmt.Sprintf("Error: invalid credentials: %v", err), http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	log.Println("Signing token with SECRET:", os.Getenv("SECRET"))
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid credentials: %v", err), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   3600 * 24 * 30,
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Login successful", "token": tokenString})
}

func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Authorization")
	if err != nil {
		http.Error(w, "No token provided", http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value
	user, err := h.userRepo.GetUserFromToken(tokenString)
	if err != nil {
		http.Error(w, "Failed to get user data", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   -1,
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}
