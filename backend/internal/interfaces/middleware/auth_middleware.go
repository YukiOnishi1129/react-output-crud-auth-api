package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey contextKey = "user"

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func AuthMiddleware(next http.Handler)http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}
		claims := &Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			http.Error(w, "Failed to parse token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims.Email)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}


func GetUserEmail(r *http.Request) string {
	email, ok := r.Context().Value(userContextKey).(string)
	if !ok {
		return ""
	}
	return email
}