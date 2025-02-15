package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

func authCorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL")) 
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS") 
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type") 

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *authHandler) RegisterAuthHandlers(r *mux.Router) {
	authRouter := r.PathPrefix(constants.AuthPath).Subrouter()
	authRouter.Use(authCorsMiddleware)

	authRouter.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	// authRouter.HandleFunc("/login", optionsLoginHandler).Methods(http.MethodPost)
	authRouter.HandleFunc("/signup", h.Signup).Methods(http.MethodPost)
	// authRouter.HandleFunc("/signup", optionsPostHandler).Methods(http.MethodPost)
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


func optionsLoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("optionsLoginHandler")
	log.Printf("Access-Control-Allow-Origin: %v", os.Getenv("FRONTEND_URL"))
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_URL"))
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}