package apis

import (
	"flicksfi/internal/config"
	"flicksfi/internal/interfaces"
	"flicksfi/internal/types"
	"flicksfi/internal/utils"
	"flicksfi/packages/movie"
	"flicksfi/packages/translate"
	"flicksfi/packages/youtube"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	service        interfaces.IApis
	userService    interfaces.IUser
	limiterService interfaces.ILimiter
}

func NewHandler(service interfaces.IApis, userService interfaces.IUser, limiterService interfaces.ILimiter) *Handler {
	return &Handler{
		service:        service,
		userService:    userService,
		limiterService: limiterService,
	}
}

func (h Handler) RegisterRoutes(router *mux.Router) {
	// Получение видео с YouTube (20)
	router.HandleFunc("/youtube/video", h.handleYouTubeVideo).Methods("GET")
	// Получаем самые популярные видео YouTube (10)
	router.HandleFunc("/youtube/video/popular", h.handlePopularYouTubeVideo).Methods("GET")
	// Получение фильмов (10) page = 0
	router.HandleFunc("/movies", h.handleMovies).Methods("GET")
	// Получение фотографии фильма
	router.HandleFunc("/image/w500/{img}", h.handleImageMovie).Methods("GET")
	// Получение фильмов по тексту (20) page = 0
	router.HandleFunc("/text/movies", h.handleMoviesText).Methods("POST")
	// Получение ТОП фильмов (20) page = 1
	router.HandleFunc("/movies/popular", h.handlePopularFilms).Methods("GET")
	// Получение информации о фильме по ID
	router.HandleFunc("/movie/{id}", h.handleMovieDetails).Methods("GET")
	// Получение похожих фильмов
	router.HandleFunc("/movies/similar", h.handleMovieSimilar).Methods("GET")
}

var total_page int = 500

// @Summary 20 videos from YouTube
// @Tags video
// @Description Getting videos from YouTube is a maximum of 20
// @ID video-youtube-20
// @Accept json
// @Produce json
// @Param categoryId query string false "Video category"
// @Param s query string false "Video Search"
// @Param y query string false "Video year"
// @Param ch query string false "Video channel"
// @Success 200 {object} []types.SearchResult
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /youtube/video [get]
func (h Handler) handleYouTubeVideo(w http.ResponseWriter, r *http.Request) {
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	query := r.URL.Query()

	// Параметры
	category := query.Get("categoryId") // ID категорий
	search := query.Get("s")            // Поиск
	year := query.Get("y")              // Год
	channel := query.Get("ch")          // Канал

	// Данные для api youtube
	var youtube_data map[string]string = map[string]string{
		"category": category,
		"search":   search,
		"year":     year,
		"channel":  channel,
	}


	videos, err := youtube.Request(youtube_data, 20)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, videos)
}

// @Summary 10 popular YouTube videos
// @Tags video
// @Description Getting 10 popular videos from YouTube
// @ID video-popular-youtube-10
// @Accept json
// @Produce json
// @Success 200 {object} []types.SearchResult
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /youtube/video/popular [get]
func (h Handler) handlePopularYouTubeVideo(w http.ResponseWriter, r *http.Request) {
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	popular, err := youtube.GetPopular()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, popular)
}

// @Summary movies by criteria
// @Tags movie
// @Description Getting 10 movies based on different criteria
// @ID movie-criteria-10
// @Accept json
// @Produce json
// @Param genre_id query string false "Movie genre"
// @Param date query string false "The year of the film's release"
// @Param s query string false "Movie Search"
// @Success 200 {object} []types.Movie
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /movies [get]
func (h Handler) handleMovies(w http.ResponseWriter, r *http.Request) {
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	query := r.URL.Query()

	var genre_id string = query.Get("genre_id")
	var release_date string = query.Get("date")
	search, err := url.QueryUnescape(query.Get("s"))
	if err != nil {
		return
	}

	var parametrs map[string]string = map[string]string{
		"genre_id":     genre_id,
		"release_date": release_date,
		"search":       search,
	}

	movies, err := movie.GetMovies(parametrs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, movies)
}

// @Summary movie poster
// @Tags movie
// @Description Getting a movie poster
// @ID movie-poster
// @Produce image/jpeg
// @Param img path string true "Movie poster"
// @Success 200 {file} string ""
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /image/w500/{img} [get]
func (h Handler) handleImageMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var img string = vars["img"]
	if img == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("path image is required"))
		return
	}

	// URL TMDB
	var url string = fmt.Sprintf("https://image.tmdb.org/t/p/w500/%s?api_key=%s", img, config.Envs.TMDB_KEY_API)

	// Запрашиваем изображение
	tmdbResponse, err := http.Get(url)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer tmdbResponse.Body.Close()

	if tmdbResponse.StatusCode != http.StatusOK {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.Header().Set("Expires", time.Now().Add(24*time.Hour).Format(http.TimeFormat))

	_, err = io.Copy(w, tmdbResponse.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

// @Summary A film based on the plot
// @Tags movie
// @Description Getting 100 movie's based on a user's story
// @ID movie-plot-100
// @Accept json
// @Produce json
// @Param DTO body types.TextMoviePayload true "Data for movies plot"
// @Success 200 {object} []types.Movie
// @Failure 400 {object} types.ErrorResponse "invalid payload"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /text/movies [post]
func (h Handler) handleMoviesText(w http.ResponseWriter, r *http.Request) {
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	var payload *types.TextMoviePayload // 1)simple 2)exact

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	text_to_en, err := translate.EN(payload.Text)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Получение фильмов по сюжету
	if payload.Lege == "simple" {
		movies, err := movie.OverviewText(text_to_en)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, movies)
		return
	}

	// Получение фильма с ИИ
	if payload.Lege == "exact" {
		// Получение текущего пользователя
		u, err := h.userService.GetUserByUUID(payload.UUID)
		if err != nil || u == nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("ip address not found"))
			return
		}

		// Получение лимитов пользователя
		limiter, err := h.limiterService.GetLimits(u.UUID)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		// Проверка лимита текста
		if limiter.TextLimiter <= 0 {
			utils.WriteError(w, http.StatusForbidden, fmt.Errorf("the limit for the day is over"))
			return
		}

		// Получение фильмов из ИИ
		movies, err := movie.GIGA_CHAT_OVERVIEW(text_to_en)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		// Уменьшение лимита
		if err := h.limiterService.ReducingLimitText(u.UUID); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusOK, movies)
		return
	}
}

// @Summary Popular movies
// @Tags movie
// @Description Getting 20 popular movies
// @ID movie-popular-20
// @Accept json
// @Produce json
// @Param page query string false "Movie Page"
// @Success 200 {object} []types.Movie
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /movies/popular [get]
func (h Handler) handlePopularFilms(w http.ResponseWriter, r *http.Request) {
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	query := r.URL.Query()
	var page string = query.Get("page")

	rand.Seed(time.Now().UnixNano())

	if page == "" {
		page = strconv.Itoa(rand.Intn(total_page) + 1)
	}

	// Конвертируем string to int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Получение популярных фильмов
	movies, err := movie.PopularMovie(pageInt)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, movies)
}

// @Summary Movie Details
// @Tags movie
// @Description Getting movie details by ID
// @ID movie-details-id
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} types.Movie
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /movie/{id} [get]
func (h Handler) handleMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	vars := mux.Vars(r)

	var id string = vars["id"]
	idINT, err := strconv.Atoi(id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Получение деталей фильма по id
	movie, err := movie.MovieDetails(idINT)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Проверка что фильм найден
	if movie.Id == 0 || movie.Title == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("movie not found"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, movie)
}

// @Summary Related Movies
// @Tags movie
// @Description Getting similar movies based on data <100
// @ID movie-similar-<100
// @Accept json
// @Produce json
// @Param genre_id query string false "The genre of the current movie"
// @Param title query string false "The title of the current movie"
// @Param overview query string false "Description of the current movie"
// @Success 200 {object} []types.Movie
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /movies/similar [get]
func (h Handler) handleMovieSimilar(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	genre_id, err := h.service.ArrayGenreIDS(query.Get("genre_id"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	title, err := url.QueryUnescape(query.Get("title"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	overview, err := url.QueryUnescape(query.Get("overview"))
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var movie_data map[string]interface{} = map[string]interface{}{
		"genre_id": genre_id,
		"title":    title,
		"overview": overview,
	}

	similar, err := movie.Similar(movie_data)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, similar)
}
