package admin

import (
	"go-user-service/internal/common/utils"
	"go-user-service/internal/packages"
	"net/http"

	"github.com/noneandundefined/vision-go"
)

type Handler struct {
	monitor *vision.Vision
	errors  *packages.Errors
}

func NewHandler(monitor *vision.Vision, errors *packages.Errors) *Handler {
	return &Handler{
		monitor: monitor,
		errors:  errors,
	}
}

func (h Handler) MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, h.monitor.GetVisionStats())
}
