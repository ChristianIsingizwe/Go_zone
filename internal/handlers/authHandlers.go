package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ChristianIsingizwe/Go_zone/internal/models"
	"github.com/ChristianIsingizwe/Go_zone/internal/services"
	"github.com/ChristianIsingizwe/Go_zone/internal/types"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Response){
	var req types.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return 
	}

	var existingUser models.User

	if err := services.DB.Where("email= ?", req.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "User already exist", http.StatusConflict)
		return
	}

	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return 
	}

	user := models.User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	if err := services.DB.Create(&user).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}