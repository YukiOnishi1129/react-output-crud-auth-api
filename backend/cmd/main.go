package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	persistence_gorm "github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/infrastructure/persistence/gorm"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/interfaces/handler"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/database"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Printf("Start server")
	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}
	r := mux.NewRouter()
	userRepository := persistence_gorm.NewUserRepository(db)
	todoRepository := persistence_gorm.NewTodoRepository(db)
	authUsecase := usecase.NewAuthUseCase(userRepository)
	todoUsecase := usecase.NewTodoUseCase(todoRepository)
	authHandler := handler.NewAuthHandler(authUsecase)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_URL")}), 
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), 
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), 
	)

	r.Use(corsOptions)

	authHandler.RegisterHandlers(r)
	todoHandler.RegisterHandlers(r)


	log.Printf("Server started at http://localhost:%s", os.Getenv("BACKEND_CONTAINER_POST"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BACKEND_CONTAINER_POST")), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}