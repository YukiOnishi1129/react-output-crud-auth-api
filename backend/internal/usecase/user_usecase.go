package usecase

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/infrastructure/persistence/dto"
	apperrors "github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/errors"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/repository"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase/input"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase/output"
)


type UserUseCase interface {
	GetUserByEmail(ctx context.Context, input *input.GetUserByEmailInput) (*output.UserOutput, error)
}

type useUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &useUseCase{userRepo: userRepo}
}


func (u *useUseCase) GetUserByEmail(ctx context.Context, input *input.GetUserByEmailInput) (*output.UserOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, apperrors.NewValidationError("invalid input parameters", err)
	}
	user, err := u.userRepo.FindByEmail(ctx, &dto.FindUserByEmailInput{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}

	return output.ConvertUserOutput(user), nil
}