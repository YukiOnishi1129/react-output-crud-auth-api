package main

import (
	"log"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/database"
	"github.com/joho/godotenv"
)


func main() {
	log.Printf("Start migrate")
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatalf("Error loading .env.local file: %v", err)
		return
	}

	db, err := database.InitConnectDB()
	if err != nil {
		log.Fatalf("Error connect to database: %v", err)
		return
	}

	// UUID拡張機能を有効化
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	db.AutoMigrate(&domain.Todo{})

	log.Printf("Migration completed")
}
