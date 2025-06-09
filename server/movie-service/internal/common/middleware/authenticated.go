package middleware

import (
	"context"
	"go-movie-service/internal/common/database/action"
	"go-movie-service/internal/common/utils"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth-token")
		if err != nil || cookie.Value == "" {
			utils.WriteJSON(w, r, http.StatusUnauthorized, "sign in to your account")
			return
		}

		uuid := cookie.Value

		user, err := action.GetUserByUUID(uuid)
		if err != nil {
			utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		if user == nil {
			utils.WriteJSON(w, r, http.StatusUnauthorized, "sign in to your account")
			return
		}

		//nolint
		ctx := context.WithValue(r.Context(), "identity", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
