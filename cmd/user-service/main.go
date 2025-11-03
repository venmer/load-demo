package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"user-service/internal/database"
	"user-service/internal/handlers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к БД
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация БД
	if err := database.InitDB(db); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Создание handlers
	userHandler := handlers.NewUserHandler(db)

	// Настройка роутера
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheck).Methods("GET")
	r.HandleFunc("/api/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/api/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("User Service starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "user-service"})
}
