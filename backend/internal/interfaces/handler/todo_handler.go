package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/interfaces/middleware"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/constants"
	apperrors "github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/errors"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase/input"
)

type TodoHandler interface {
	RegisterTodoHandlers(r *mux.Router)
	ListTodo(w http.ResponseWriter, r *http.Request)
	GetTodo(w http.ResponseWriter, r *http.Request)
	CreateTodo(w http.ResponseWriter, r *http.Request)
	UpdateTodo(w http.ResponseWriter, r *http.Request)
	DeleteTodo(w http.ResponseWriter, r *http.Request)
}
type todoHandler struct {
	BaseHandler
	todoUseCase usecase.TodoUseCase
	userUseCase usecase.UserUseCase
}

func NewTodoHandler(todoUseCase usecase.TodoUseCase, userUseCase usecase.UserUseCase) TodoHandler {
	return &todoHandler{todoUseCase: todoUseCase, userUseCase: userUseCase}
}


func (h *todoHandler) RegisterTodoHandlers(r *mux.Router) {
	todoRouter := r.PathPrefix(constants.TodosPath).Subrouter()
	todoRouter.Use(middleware.AuthMiddleware)

	todoRouter.HandleFunc("", h.ListTodo).Methods("GET")
	todoRouter.HandleFunc("/{id}", h.GetTodo).Methods("GET")
	todoRouter.HandleFunc("", h.CreateTodo).Methods("POST")
	todoRouter.HandleFunc("", optionsPostHandler).Methods("OPTIONS")
	todoRouter.HandleFunc("/{id}", h.UpdateTodo).Methods("PUT")
	todoRouter.HandleFunc("/{id}", h.DeleteTodo).Methods("DELETE")
	todoRouter.HandleFunc("/{id}", optionsDeleteHandler).Methods("OPTIONS")
}

func (h *todoHandler) ListTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	email := middleware.GetUserEmail(r)
	user, err := h.userUseCase.GetUserByEmail(ctx, &input.GetUserByEmailInput{Email: email})
	if err != nil {
		h.respondError(w, err)
		return
	}

	output, err := h.todoUseCase.ListTodo(ctx, &input.ListTodoInput{UserID: user.ID})
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}

func (h *todoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
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

func (h *todoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
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

func (h *todoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
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

func (h *todoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
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



