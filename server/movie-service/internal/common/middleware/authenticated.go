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
		if err != nil {
			utils.WriteJSON(w, r, http.StatusUnauthorized, "Sign in to your account")
			return
		}

		uuid, err := utils.Decrypt(cookie.Value)
		if err != nil {
			utils.WriteJSON(w, r, http.StatusInternalServerError, err)
			return
		}

		user, err := action.GetUserByUUID(uuid)
		if err != nil {
			utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		if user == nil {
			utils.WriteJSON(w, r, http.StatusUnauthorized, "Sign in to your account")
			return
		}

		//nolint
		ctx := context.WithValue(r.Context(), "identity", user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// func GetUserFromContext(ctx context.Context) *User {
// 	user, ok := ctx.Value(UserKey).(*User)
// 	if !ok {
// 		return nil // Или panic, если это критическая ошибка
// 	}
// 	return user
// }
