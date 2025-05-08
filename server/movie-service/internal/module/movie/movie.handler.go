package movie

import (
	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/lib/movie"
	"go-movie-service/internal/packages"
	"net/http"

	"github.com/noneandundefined/vision-go"
)

type Handler struct {
	monitor *vision.Vision
	errors  *packages.Errors
	movie   *movie.Movie
}

func NewHandler(monitor *vision.Vision, errors *packages.Errors) *Handler {
	return &Handler{
		monitor: monitor,
		errors:  errors,
		movie:   movie.NewMovie(),
	}
}

func (h Handler) MovieGetHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := h.movie.GetMovies()
	if err != nil {
		utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, r, http.StatusOK, movies)
}
