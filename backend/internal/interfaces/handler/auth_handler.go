package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/constants"
	apperrors "github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/errors"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase"
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/usecase/input"
	"github.com/gorilla/mux"
)



type AuthHandler interface {
	RegisterAuthHandlers(r *mux.Router)
	Login(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
}

type authHandler struct {
	BaseHandler
	authUseCase usecase.AuthUseCase
}

func NewAuthHandler(authUseCase usecase.AuthUseCase) AuthHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (h *authHandler) RegisterAuthHandlers(r *mux.Router) {
	authRouter := r.PathPrefix(constants.AuthPath).Subrouter()

	authRouter.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	authRouter.HandleFunc("/login", optionsPostHandler).Methods(http.MethodPost)
	authRouter.HandleFunc("/signup", h.Signup).Methods(http.MethodPost)
	authRouter.HandleFunc("/signup", optionsPostHandler).Methods(http.MethodPost)
}

func (h *authHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := &input.LoginInput{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}

	output, err := h.authUseCase.Login(ctx, input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusOK, output)
}


func (h *authHandler) Signup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := &input.RegisterUserInput{}
	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		h.respondError(w, apperrors.NewValidationError("invalid request body", err))
		return
	}

	output, err := h.authUseCase.RegisterUserInput(ctx, input)
	if err != nil {
		h.respondError(w, err)
		return
	}

	h.respondJSON(w, http.StatusCreated, output)
}