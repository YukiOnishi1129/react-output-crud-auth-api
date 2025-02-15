package handler

import (
	"net/http"

	"github.com/YukiOnishi1129/react-output-crud-auth-api/backend/internal/interfaces/middleware"
)

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	email := middleware.GetUserEmail(r)
	if email == "" {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// ユーザー情報を使用して処理を行う
	w.Write([]byte("Welcome, " + email))
}