package handler

import (
	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/pkg/constants"
	"github.com/gorilla/mux"
)

func (h *TodoHandler) RegisterHandlers(r *mux.Router) {
	todoRouter := r.PathPrefix(constants.TodosPath).Subrouter()

	todoRouter.HandleFunc("", h.ListTodo).Methods("GET")
	todoRouter.HandleFunc("/{id}", h.GetTodo).Methods("GET")
	todoRouter.HandleFunc("", h.CreateTodo).Methods("POST")
	todoRouter.HandleFunc("", optionsPostHandler).Methods("OPTIONS")
	todoRouter.HandleFunc("/{id}", h.UpdateTodo).Methods("PUT")
	todoRouter.HandleFunc("/{id}", h.DeleteTodo).Methods("DELETE")
	todoRouter.HandleFunc("/{id}", optionsDeleteHandler).Methods("OPTIONS")
}