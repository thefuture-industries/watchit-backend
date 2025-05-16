package movie

import (
	"fmt"
	"net/http"
	"strconv"

	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/packages"

	"github.com/gorilla/mux"
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

func (h Handler) MovieDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		utils.WriteJSON(w, r, http.StatusBadRequest, "movie ID is required")
		return
	}

	idINT, err := strconv.Atoi(id)
	if err != nil {
		packages.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid movie ID format"))
		return
	}

	movie, err := MovieDetails(idINT)
	if err != nil {
		packages.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if movie.Id == 0 || movie.Title == "" {
		packages.WriteError(w, http.StatusNotFound, fmt.Errorf("movie not found"))
		return
	}

	if err := packages.CacheJSON(w, 3600, http.StatusOK, movie); err != nil {
		packages.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error writing response: %w", err))
		return
	}
}
