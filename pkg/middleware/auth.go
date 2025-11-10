package middleware

import (
	"context"
	"net/http"
	"strings"

	"task_API/internal/storage"
	"task_API/pkg/utils"
)

func AuthMiddleware(storage *storage.PostgresStorage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			// Auth Header
			authHeader := request.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(writer, "Authorization header is required", http.StatusUnauthorized)
				return
			}

			// Bearer check
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				http.Error(writer, "Invalid Authorization header format", http.StatusUnauthorized)
				return
			}

			// Token Validation
			tokenStr := tokenParts[1]
			claim, error := utils.ValidateJWTToken(tokenStr)
			if error != nil {
				http.Error(writer, "Invalid token: "+error.Error(), http.StatusUnauthorized)
				return
			}

			// Get User from DB
			user, error := storage.GetUserByID(claim.UserId)
			if error != nil {
				http.Error(writer, "User not found: "+error.Error(), http.StatusUnauthorized)
				return
			}

			// Add user to context
			ctx := context.WithValue(request.Context(), "user", user)
			next.ServeHTTP(writer, request.WithContext(ctx))
		})
	}
}
