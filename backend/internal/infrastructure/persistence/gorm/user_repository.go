package persistence_gorm

import (
	"context"

	"gorm.io/gorm"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/infrastructure/persistence/dto"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/repository"
)


type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}


func (r *userRepository) FindByEmail(ctx context.Context, input *dto.FindUserByEmailInput) (*dto.UserOutput, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", input.Email).Error; err != nil {
		return nil, HandleDBError(err, "user")
	}
	return dto.ConvertUserOutput(&user), nil
}

func (r *userRepository) Create(ctx context.Context, input *dto.CreateUserInput) (*dto.UserOutput, error) {
	user := domain.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return nil, HandleDBError(err, "user")
	}
	return dto.ConvertUserOutput(&user), nil
}