package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	apperrors "github.com/YukiOnishi1129/react-output-crud-api/backend/internal/pkg/errors"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase"
	"github.com/YukiOnishi1129/react-output-crud-api/backend/internal/usecase/input"
)

type TodoHandlerInterface interface {
	RegisterHandlers(r *mux.Router)
	ListTodo(w http.ResponseWriter, r *http.Request)
	GetTodo(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}


type TodoHandler struct {
	BaseHandler
	todoUseCase usecase.TodoUseCase
}

func NewTodoHandler(todoUseCase usecase.TodoUseCase) TodoHandlerInterface {
	return &TodoHandler{todoUseCase: todoUseCase}
}

func (h *TodoHandler) ListTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	output, err := h.todoUseCase.ListTodo(ctx)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid todo id", err))
		return
	}

	input := &input.GetTodoInput{
		ID:     todoID,
	}

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	output, err := h.todoUseCase.GetTodo(ctx, input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var input input.CreateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	output, err := h.todoUseCase.CreateTodo(ctx, &input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusCreated, output)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid todo id", err))
		return
	}

	var input input.UpdateTodoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}
	input.ID = todoID

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	output, err := h.todoUseCase.UpdateTodo(ctx, &input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	todoID, err := uuid.Parse(vars["id"])
	if err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid todo id", err))
		return
	}

	input := &input.DeleteTodoInput{
		ID:     todoID,
	}

	if err := input.Validate(); err != nil {
		h.respondError(w, apperrors.NewValidationError("validation failed", err))
		return
	}

	if err := h.todoUseCase.DeleteTodo(ctx, input); err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusNoContent, nil)
}