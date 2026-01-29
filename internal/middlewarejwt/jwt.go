package middlewarejwt

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserIDKey   contextKey = "user_id"
	ContextTenantIDKey contextKey = "tenant_id"
	ContextRoleIDKey   contextKey = "role_id"
	ContextModulesKey  contextKey = "modules"
)

type AuthClaims struct {
	UserID            uint              `json:"sub"`
	TenantID          uint              `json:"tenant"`
	RoleID            uint              `json:"role_id"`
	ModulePermissions map[uint][]string `json:"modules"`
	jwt.RegisteredClaims
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))

		claims := &AuthClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		if claims.ModulePermissions == nil {
			claims.ModulePermissions = map[uint][]string{}
		}

		ctx := context.WithValue(r.Context(), ContextUserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, ContextTenantIDKey, claims.TenantID)
		ctx = context.WithValue(ctx, ContextRoleIDKey, claims.RoleID)
		ctx = context.WithValue(ctx, ContextModulesKey, claims.ModulePermissions)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
