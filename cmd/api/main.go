package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ChristianIsingizwe/Go_zone/configs"
	"github.com/ChristianIsingizwe/Go_zone/internal/handlers"
	"github.com/ChristianIsingizwe/Go_zone/internal/services"
)

func main() {

	configs.LoadEnv()

	if err := services.ConnectToDatabase(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database...")

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.RegisterHandler(w, r)
		} else{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			handlers.LoginHandler(w, r)
		}else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	port := os.Getenv("APP_PORT")
	if port == ""{
		port = ":8080"
	}
	log.Printf("Starting server on port %s...\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

}
