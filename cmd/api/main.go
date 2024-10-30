package main

import (
	"net/http"

	"github.com/ChristianIsingizwe/Go_zone/internal/handlers"
	"github.com/ChristianIsingizwe/Go_zone/internal/services"
)

func main() {
	services.ConnectToDatabase()

	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/refresh", handlers.RefreshTokenHandler)
}