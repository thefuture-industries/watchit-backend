package movie

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

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

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "not all fields are filled in!")
		return
	}

	movies, err := h.movie.GetMoviesByText(payload.Text)
	if err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, r, http.StatusOK, movies)
}

func (h Handler) MovieTextFREEHandler(w http.ResponseWriter, r *http.Request) {
	var payload *types.TMoviesPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "not all fields are filled in!")
		return
	}

	movies, err := h.movie.GetMoviesByText(payload.Text)
	if err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSON(w, r, http.StatusOK, movies)
}

func (h Handler) MovieImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	image, ok := vars["image"]

	if !ok {
		utils.WriteJSON(w, r, http.StatusBadRequest, "image not found")
		return
	}

	url := fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s?api_key=%s", image, os.Getenv("TMDB_KEY_API"))
	fmt.Println(url)

	httpConfig, err := utils.GetProxy()
	if err != nil {
		utils.WriteJSON(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	client := &http.Client{
		Transport: httpConfig,
		Timeout:   15 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		utils.WriteJSON(w, r, http.StatusBadGateway, err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println(resp)

	if resp.StatusCode != http.StatusOK {
		utils.WriteJSON(w, r, resp.StatusCode, "error send request to get image")
		return
	}

	w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	w.WriteHeader(http.StatusOK)

	io.Copy(w, resp.Body)
}
