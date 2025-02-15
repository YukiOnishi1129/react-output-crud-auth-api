package main

import (
	"log"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/database"
	"github.com/joho/godotenv"
)


func main() {
	log.Printf("Start rollback")
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

	err = db.Migrator().DropTable(&domain.Todo{})
	if err != nil {
		log.Fatalf("Error dropping tables: %v", err)
		return
	}

	log.Printf("Dropped tables")
}
