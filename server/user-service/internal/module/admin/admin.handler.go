package admin

import (
	"go-user-service/cmd/conf"
	"go-user-service/internal/common/utils"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	monitor *conf.Vision
	logger  *zap.Logger
}

func NewHandler(monitor *conf.Vision, logger *zap.Logger) *Handler {
	return &Handler{
		monitor: monitor,
		logger:  logger,
	}
}

func (h Handler) MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, h.monitor.GetVisionStats())
}
