package admin

import (
	"go-user-service/internal/packages"
	"go-user-service/internal/common/utils"
	"net/http"

	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
)

type Handler struct {
	monitor *vision.Vision
	logger  *zap.Logger
	errors  *packages.Errors
}

func NewHandler(monitor *vision.Vision, logger *zap.Logger, errors *packages.Errors) *Handler {
	return &Handler{
		monitor: monitor,
		logger:  logger,
		errors:  errors,
	}
}

func (h Handler) MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, h.monitor.GetVisionStats())
}
