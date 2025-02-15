package usecase

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/infrastructure/persistence/dto"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/auth"
	apperrors "github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/errors"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/repository"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase/input"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase/output"
)



type AuthUseCase interface {
	Login(ctx context.Context, input *input.LoginInput) (*output.AuthOutput, error)
	RegisterUserInput(ctx context.Context, input *input.RegisterUserInput) (*output.AuthOutput, error)
}

type authUseCase struct {
	userRepo repository.UserRepository
}

func NewAuthUseCase(userRepo repository.UserRepository) AuthUseCase {
	return &authUseCase{userRepo: userRepo}
}


func (u *authUseCase) Login(ctx context.Context, input *input.LoginInput) (*output.AuthOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, apperrors.NewValidationError("invalid input parameters", err)
	}
	// find user by email
	user, err := u.userRepo.FindByEmail(ctx, &dto.FindUserByEmailInput{
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}

	// verify password
	if err :=  auth.VerifyPassword(user.Password, input.Password); err != nil {
		return nil, apperrors.NewUnauthorizedError("email or password is incorrect", err)
	}

	// create jwt token
	tokenString, err := auth.GenerateToken(user.Email)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to create token", err)
	}

	userOutput := output.ConvertUserOutput(user)
	return &output.AuthOutput{
		Token: tokenString,
		User:  *userOutput,
	}, nil
}


func (u *authUseCase) RegisterUserInput(ctx context.Context, input *input.RegisterUserInput) (*output.AuthOutput, error) {
	if err := input.Validate(); err != nil {
		return nil, apperrors.NewValidationError("invalid input parameters", err)
	}

	// hash password
	hashedPassword, err := auth.HashPassword(input.Password)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to hash password", err)
	}

	// create user
	user, err := u.userRepo.Create(ctx, &dto.CreateUserInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	// create jwt token
	tokenString, err := auth.GenerateToken(user.Email)
	if err != nil {
		return nil, apperrors.NewInternalError("failed to create token", err)
	}

	userOutput := output.ConvertUserOutput(user)
	return &output.AuthOutput{
		Token: tokenString,
		User:  *userOutput,
	}, nil
}