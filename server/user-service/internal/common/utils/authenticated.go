package utils

import (
	"context"
	"fmt"
	"go-user-service/internal/common/database/actions"
	"net/http"
)

// UserContextKey - тип для ключа контекста пользователя
type UserContextKey string

const UserKey UserContextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth-token")
		if err != nil {
			WriteJSON(w, r, http.StatusUnauthorized, fmt.Errorf("Sign in to your account"))
			return
		}

		uuid, err := Decrypt(cookie.Value)
		if err != nil {
			WriteJSON(w, r, http.StatusInternalServerError, err)
			return
		}

		user, err := actions.GetUserByUUID(uuid)
		if err != nil {
			WriteJSON(w, r, http.StatusInternalServerError, err.Error())
			return
		}

		if user == nil {
			WriteJSON(w, r, http.StatusUnauthorized, fmt.Errorf("Sign in to your account"))
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, uuid)
		r = r.WithContext(ctx)

		// Cookie найден - продолжаем обработку
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
