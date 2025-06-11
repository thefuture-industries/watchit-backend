package favourite

import (
	"go-movie-service/internal/packages"
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

func (h *Handler) GetFavouritesHandler(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) AddFavouriteHandler(w http.ResponseWriter, r *http.Request) {}
