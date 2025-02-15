package domain

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"` 
	Title  string    `json:"title"`
	Content *string    `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
}

func (Todo) TableName() string {
	return "todos"
}