package admin

import (
	"go-user-service/internal/common/utils"
	"net/http"

	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
)

type Handler struct {
	monitor *vision.Vision
	logger  *zap.Logger
}

func NewHandler(monitor *vision.Vision, logger *zap.Logger) *Handler {
	return &Handler{
		monitor: monitor,
		logger:  logger,
	}
}

func (h Handler) MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, h.monitor.GetVisionStats())
}
