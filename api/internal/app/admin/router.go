package admin

import (
	"flicksfi/internal/interfaces"
	"flicksfi/internal/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	service     interfaces.IAdmin
	userService interfaces.IUser
}

func NewHandler(service interfaces.IAdmin, userService interfaces.IUser) *Handler {
	return &Handler{
		service:     service,
		userService: userService,
	}
}

func (h Handler) RegisterRoutes(router *mux.Router) {
	// Получение статистики приложения
	router.HandleFunc("/statistic", h.handleGetStatistic).Methods("GET")
	router.HandleFunc("/monitoring", h.handleGetMonitoring).Methods("GET")
}

// -------------------------------
// Получение статистики приложения
// -------------------------------
func (h Handler) handleGetStatistic(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// users, err := h.userService.GetUsers()
	// if err != nil {
	// 	utils.WriteError(w, http.StatusInternalServerError, err)
	// 	return
	// }
}

// --------------------------------
// Получение мониторинга приложения
// --------------------------------
func (h Handler) handleGetMonitoring(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, h.service.GetMonitoring())
}
