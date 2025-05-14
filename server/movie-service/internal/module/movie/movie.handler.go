package movie

import (
	"fmt"
	"go-movie-service/internal/packages"
	"net/http"
	"strconv"

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

// MovieDetailsHandler обрабатывает запрос на получение деталей фильма по ID
// @Summary Movie Details
// @Tags movie
// @Description Getting movie details by ID
// @ID movie-details-id
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} types.Movie
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 404 {object} types.ErrorResponse "Movie Not Found"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /movie/{id} [get]
func (h Handler) MovieDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		packages.WriteError(w, http.StatusBadRequest, fmt.Errorf("movie ID is required"))
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
