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
		utils.WriteJSON(w, r, http.StatusBadRequest, "couldn't find the movie by ID")
		return
	}

	idINT, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "movie ID conversion error")
		return
	}

	movie, err := MovieDetails(idINT)
	if err != nil {
		utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	if movie.Id == 0 || movie.Title == "" {
		utils.WriteJSON(w, r, http.StatusNotFound, fmt.Errorf("we didn't find any movies with id: %d", idINT))
		return
	}

	utils.CacheJSON(w, 1800) // 30 min.
	utils.WriteJSON(w, r, http.StatusOK, movie)
}
