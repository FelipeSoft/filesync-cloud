package middleware

import (
	"net/http"

	"github.com/FelipeSoft/filesync-cloud/internal/domain"
	httputil "github.com/FelipeSoft/filesync-cloud/internal/utils/http"
)

type AuthMiddleware struct {
	tokenManager domain.TokenManager
}

func NewAuthMiddleware(tokenManager domain.TokenManager) *AuthMiddleware {
	return &AuthMiddleware{
		tokenManager: tokenManager,
	}
}

func (m *AuthMiddleware) Handle(w http.ResponseWriter, r *http.Request, next func(w http.ResponseWriter, r *http.Request)) {
	token := r.Header.Get("Authorization")

	if token == "" {
		httputil.WriteJSON(w, http.StatusUnauthorized, httputil.HttpResponse{
			Error: "Could not proceed without Bearer Token.",
		})
		return
	}
	next(w, r)
}
