package persistence_gorm

import (
	"context"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/domain"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/infrastructure/persistence/dto"
	apperrors "github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/errors"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/repository"
	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{db: db}
}

func (r *todoRepository) FindAll(ctx context.Context, input *dto.FindAllInput) (*dto.TodoListOutput, error) {
	var todos []*domain.Todo
	if err := r.db.Where("user_id = ?", input.UserID).Find(&todos).Error; err != nil {
		return &dto.TodoListOutput{}, err
	}
	return dto.ConvertTodoListOutput(todos, int64(len(todos))), nil
}

func (r *todoRepository) FindByID(ctx context.Context, input *dto.FindByIDInput) (*dto.TodoOutput, error){
	var todo domain.Todo
	if err := r.db.First(&todo, "id = ?", input.ID).Error; err != nil {
		return nil, HandleDBError(err, "todo")
	}
	
	return dto.ConvertTodoOutput(&todo), nil
}

func (r *todoRepository) Create(ctx context.Context, input *dto.CreateTodoInput) (*dto.TodoOutput, error) {
	var todo domain.Todo
	todo.Title = input.Title
	todo.Content = input.Content
	if err := r.db.Create(&todo).Error; err != nil {
		return nil, HandleDBError(err, "todo")
	}
	return dto.ConvertTodoOutput(&todo), nil
}

func (r *todoRepository) Update(ctx context.Context, input *dto.UpdateTodoInput) (*dto.TodoOutput, error){
	var todo domain.Todo
	todo.ID = input.ID
	todo.Title = input.Title
	todo.Content = input.Content
	if err := r.db.Save(&todo).Error; err != nil {
		return nil, HandleDBError(err, "todo")
	}
	return dto.ConvertTodoOutput(&todo), nil
}

func (r *todoRepository) Delete(ctx context.Context, input *dto.DeleteTodoInput) error {
	result := r.db.Delete(&domain.Todo{}, "id = ?", input.ID)
	if result.Error != nil {
		return HandleDBError(result.Error, "todo")
	}
	if result.RowsAffected == 0 {
		return apperrors.NewNotFoundError("todo not found", result.Error)
	}
	return nil
}