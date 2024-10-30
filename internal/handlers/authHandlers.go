package handlers

import (
	"encoding/json"
	"fmt"
	"go/token"
	"net/http"

	"github.com/ChristianIsingizwe/Go_zone/internal/models"
	"github.com/ChristianIsingizwe/Go_zone/internal/services"
	"github.com/ChristianIsingizwe/Go_zone/internal/types"
	"github.com/dgrijalva/jwt-go"
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


func LoginHandler(w http.ResponseWriter, r *http.Request){
	var req types.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return 
	}

	var user models.User

	if err := services.DB.Where("email = ? ", req.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !services.CheckPassword(req.Password, user.Password){
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return 
	}

	accessToken, err := services.GenerateAccessToken(fmt.Sprint(user.ID))
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return 
	}

	refreshToken, err := services.GenerateRefreshToken(fmt.Sprint(user.ID))
	if err != nil {
		http.Error(w, "Failed to generate  refresh token", http.StatusInternalServerError )
		return 
	}

	json.NewEncoder(w).Encode(types.TokenResponse{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	})
}


func RefreshTokenHandler(w http.ResponseWriter, r *http.Request){
	refreshToken := r.Header.Get("Authorization")
	if refreshToken == "" {
		http.Error(w, "Refresh token required", http.StatusUnauthorized)
		return 
	}

	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHS256); !ok {
			return nil, http.ErrAbortHandler
		}
		
	})
}