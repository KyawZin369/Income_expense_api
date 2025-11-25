package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/example/graphql-mysql-api/pkg/utils"
)

const UserIDKey = "user_id"
const UserEmailKey = "user_email"

// AuthMiddleware is a GraphQL middleware that validates JWT tokens
func AuthMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	// Get the operation context
	oc := graphql.GetOperationContext(ctx)
	if oc == nil {
		return next(ctx)
	}

	// Get the request headers
	headers := oc.Headers

	// Extract token from Authorization header
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		// No token provided, continue without authentication
		// Individual resolvers can check if user is authenticated
		return next(ctx)
	}

	// Extract token from "Bearer <token>" format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return next(ctx)
	}

	token := parts[1]

	// Validate token
	claims, err := utils.ValidateToken(token)
	if err != nil {
		// Invalid token, continue without authentication
		return next(ctx)
	}

	// Add user info to context
	ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
	ctx = context.WithValue(ctx, UserEmailKey, claims.Email)

	return next(ctx)
}

// GetUserIDFromContext extracts user ID from context
func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(UserIDKey).(uint)
	return userID, ok
}

// GetUserEmailFromContext extracts user email from context
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(UserEmailKey).(string)
	return email, ok
}

// RequireAuth is a helper that checks if user is authenticated
func RequireAuth(ctx context.Context) (uint, error) {
	userID, ok := GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("authentication required")
	}
	return userID, nil
}

