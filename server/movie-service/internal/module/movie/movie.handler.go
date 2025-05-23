package movie

import (
	"net/http"
	"strconv"

	"go-movie-service/internal/common/database/schema"
	"go-movie-service/internal/common/utils"
	"go-movie-service/internal/lib/movie"
	"go-movie-service/internal/packages"
	"go-movie-service/internal/types"

	"github.com/gorilla/mux"
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

	movie, err := h.movie.GetDetailsMovies(uint32(idINT))
	if err != nil {
		utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	utils.CacheJSON(w, 1800) // 30 min.
	utils.WriteJSON(w, r, http.StatusOK, movie)
}

func (h Handler) MovieGetHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := h.movie.GetMovies()
	if err != nil {
		utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSON(w, r, http.StatusOK, movies)
}

func (h Handler) MovieTextHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pUUID, okUUID := vars["uuid"]
	if !okUUID {
		utils.WriteJSON(w, r, http.StatusBadRequest, "couldn't find the user by uuid.")
		return
	}

	user := r.Context().Value("identity").(*schema.Users)

	if pUUID != user.UUID {
		utils.WriteJSON(w, r, http.StatusBadRequest, "the uuid was transmitted incorrectly.")
		return
	}

	var payload *types.TMoviesPayload

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "тot all fields are filled in")
		return
	}
}
