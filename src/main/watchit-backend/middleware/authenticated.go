// nolint
package middleware

import (
	"context"
	"net/http"
	"strings"
	"watchit/httpx/encryption"
	"watchit/httpx/pkg/httpx"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// IsAuthenticated Middleware для проверки на аутентифицированного пользователя
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			httpx.HttpResponse(w, r, http.StatusUnauthorized, "log in to your account")
			return
		}

		const prefix = "Bearer "
		if !strings.HasPrefix(header, prefix) {
			httpx.HttpResponse(w, r, http.StatusUnauthorized, "log in to your account")
			return
		}

		token := strings.TrimPrefix(header, prefix)

		authToken, err := encryption.Decrypt(token)
		if err != nil {
			httpx.HttpResponse(w, r, http.StatusUnauthorized, "log in to your account")
			return
		}

		//nolint
		ctx := context.WithValue(r.Context(), "identity", authToken)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
