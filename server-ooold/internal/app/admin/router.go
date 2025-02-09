package admin

import (
	"flicksfi/internal/config"
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

// @Summary Monitoring app
// @Tags admin
// @Description Getting application monitoring - errors, average response time of the server and database
// @ID admin-monitoring
// @Accept json
// @Produce json
// @Param key query string true "The key to access monitoring data"
// @Success 200 {object} types.MonitoringResponse
// @Failure 403 {object} types.ErrorResponse "invalid key"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /monitoring [get]
func (h Handler) handleGetMonitoring(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("key")

	if key != config.Envs.ADMIN_KEY {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("invalid key"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, h.service.GetMonitoring())
}
