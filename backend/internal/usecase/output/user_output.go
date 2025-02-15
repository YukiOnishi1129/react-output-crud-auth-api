package output

import (
	"time"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/infrastructure/persistence/dto"
	"github.com/google/uuid"
)

type UserOutput struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func ConvertUserOutput(user *dto.UserOutput) *UserOutput {
	return &UserOutput{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}