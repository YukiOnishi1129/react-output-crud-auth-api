package main

import (
	"log"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/database"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/pointer"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)


func main() {
	log.Printf("Start seed")
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

	todoID1 := uuid.New()
	todoID2 := uuid.New()

	insertTodoList := []*domain.Todo{
		{
			ID: todoID1,
			Title: "title1",
			Content: pointer.String("content1"),
		},
		{
			ID: todoID2,
			Title: "title2",
			Content: pointer.String("content2"),
		},
	}

	db.Create(insertTodoList)

	todos := []*domain.Todo{}

	db.Find(&todos)

	log.Printf("Successfully inserted %d todos", len(todos))
}
