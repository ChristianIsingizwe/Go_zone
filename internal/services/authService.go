package services

import (
	"errors"
	"go/token"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const jwtSecret = "secretKey"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "",err
	}

	return string(bytes), nil
}

func CheckPassword(hashedPassword, password string)error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func GenerateToken(userID uint)(string, error){
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}


func ParseToken(tokenStr string) (uint, error){
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("Invalid signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Invalid claims")
	}

	return uint(claims["sub"].(float64)), nil
}