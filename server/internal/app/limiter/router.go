package limiter

import (
	"flicksfi/internal/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service     interfaces.ILimiter
	userService interfaces.IUser
}

func NewHandler(service interfaces.ILimiter, userService interfaces.IUser) *Handler {
	return &Handler{
		service:     service,
		userService: userService,
	}
}

func (h Handler) RegisterRoutes(router *mux.Router) {
	// Получение лимита пользователя
	router.HandleFunc("/limiter", h.handleGetLimiter).Methods("GET")
}

func (h Handler) handleGetLimiter(w http.ResponseWriter, r *http.Request) {}
