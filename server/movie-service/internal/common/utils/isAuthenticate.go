// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package utils

// import (
// 	"context"
// 	"fmt"
// 	"go-movie-service/internal/common/database/actions"
// 	"net/http"
// )

// UserContextKey - тип для ключа контекста пользователя
type UserContextKey string

const UserKey UserContextKey = "user"

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cookie, err := r.Cookie("auth-token")
// 		if err != nil {
// 			WriteError(w, http.StatusUnauthorized, fmt.Errorf("log in to your account"))
// 			return
// 		}

// 		uuid, err := Decrypt(cookie.Value)
// 		if err != nil {
// 			WriteError(w, http.StatusInternalServerError, err)
// 			return
// 		}

// 		user, err := actions.GetUserByUUID(uuid)
// 		if err != nil {
// 			WriteError(w, http.StatusInternalServerError, err)
// 			return
// 		}

// 		if user == nil {
// 			WriteError(w, http.StatusUnauthorized, fmt.Errorf("log in to your account"))
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), UserKey, uuid)
// 		r = r.WithContext(ctx)

// 		// Cookie найден - продолжаем обработку
// 		next.ServeHTTP(w, r)
// 	})
// }

// func GetUserFromContext(ctx context.Context) *User {
// 	user, ok := ctx.Value(UserKey).(*User)
// 	if !ok {
// 		return nil // Или panic, если это критическая ошибка
// 	}
// 	return user
// }
