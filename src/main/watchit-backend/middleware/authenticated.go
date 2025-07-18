// nolint
package middleware

import (
	"context"
	"net/http"
	"time"
	"watchit/httpx/encryption"
	"watchit/httpx/infra/types"
	"watchit/httpx/pkg/httpx"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Middleware для проверки на аутентифицированного пользователя
func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth-token")
		if err != nil || cookie == nil {
			httpx.HttpResponse(w, r, http.StatusUnauthorized, "Войдите в аккаунт.")
			return
		}

		authTokenModel, err := encryption.Decrypt(cookie.Value)
		if err != nil {
			httpx.HttpResponse(w, r, http.StatusUnauthorized, "Войдите в аккаунт.")
			return
		}

		var authToken types.AuthToken
		json.Unmarshal([]byte(authTokenModel), &authToken)

		if time.Since(authToken.Timestamp).Hours() >= 24 {
			httpx.HttpResponse(w, r, http.StatusUnauthorized, "Войдите в аккаунт.")
			return
		}

		//nolint
		ctx := context.WithValue(r.Context(), "identity", authToken)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
