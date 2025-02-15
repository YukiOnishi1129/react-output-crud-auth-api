package repository

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/infrastructure/persistence/dto"
)

type TodoRepository interface {
	FindAll(ctx context.Context) (*dto.TodoListOutput, error)
	FindByID(ctx context.Context, input *dto.FindByIDInput) (*dto.TodoOutput, error)
	Create(ctx context.Context, input *dto.CreateTodoInput) (*dto.TodoOutput, error)
	Update(ctx context.Context, input *dto.UpdateTodoInput) (*dto.TodoOutput, error)
	Delete(ctx context.Context, input *dto.DeleteTodoInput) error
}