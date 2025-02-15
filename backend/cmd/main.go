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
	todoRepository := persistence_gorm.NewTodoRepository(db)
	todoUsecase := usecase.NewTodoUseCase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{os.Getenv("FRONTEND_URL")}), // 環境変数からオリジンを取得
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // 許可するHTTPメソッド
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // 許可するヘッダー
	)

	// CORSミドルウェアをルーターに適用
	r.Use(corsOptions)

	todoHandler.RegisterHandlers(r)


	log.Printf("Server started at http://localhost:%s", os.Getenv("BACKEND_CONTAINER_POST"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("BACKEND_CONTAINER_POST")), r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}