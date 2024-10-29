package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


var accessTokenSecretKey= os.Getenv("ACCESS_TOKEN_SECRET_KEY")
var refreshTokenSecretKey= os.Getenv("REFRESH_TOKEN_SECRET_KEY")



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


func GenerateAccessToken(userID string)(string, error){
	claims := jwt.MapClaims{
		"userId": userID,
		"exp": time.Now().Add(30 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessTokenSecretKey))
}

func GenerateRefreshToken(userID string) (string, error){
	claims := jwt.MapClaims{
		"userId": userID,
		"exp": time.Now().Add(7*24*time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshTokenSecretKey))
}


// func ParseToken(tokenStr string) (uint, error){
// 	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
// 			return nil, errors.New("Invalid signing method")
// 		}
// 		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
// 	})

// 	if err != nil || !token.Valid {
// 		return 0, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return 0, errors.New("Invalid claims")
// 	}

// 	return uint(claims["sub"].(float64)), nil
// }