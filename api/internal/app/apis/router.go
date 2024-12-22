package apis

import (
	"flicksfi/internal/config"
	"flicksfi/internal/interfaces"
	"flicksfi/internal/types"
	"flicksfi/internal/utils"
	"flicksfi/pkg/movie"
	"flicksfi/pkg/translate"
	"flicksfi/pkg/youtube"
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

// Получение видео с YouTube (20)
func (h Handler) handleYouTubeVideo(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Инициализация переменной для получения из URL параметров
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

	// Получаем видео
	videos, err := youtube.Request(youtube_data, 20)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, videos)
}

// Получаем самые популярные видео YouTube (10)
func (h Handler) handlePopularYouTubeVideo(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
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

// Получение фильмов (10) page = 0
func (h Handler) handleMovies(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Инициализация переменной для получения из URL параметров
	query := r.URL.Query()

	var genre_id string = query.Get("genre_id")
	var release_date string = query.Get("date")
	search, err := url.QueryUnescape(query.Get("s"))
	if err != nil {
		return
	}
	fmt.Println(search)

	// Создания данных для получения фильмов
	var parametrs map[string]string = map[string]string{
		"genre_id":     genre_id,
		"release_date": release_date,
		"search":       search,
	}

	// Получение фильмов
	movies, err := movie.GetMovies(parametrs)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, movies)
}

// Получение фотографии фильма
func (h Handler) handleImageMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Получаем img
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

// Получение фильмов по тексту (20) page = 0
func (h Handler) handleMoviesText(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.TextMoviePayload // 1)simple 2)exact

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Перевод текста на английский
	text_to_en, err := translate.EN(payload.Text)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(text_to_en)

	// Получение фильмов по сюжету
	if payload.Lege == "simple" {
		movies, err := movie.OverviewText(text_to_en)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		fmt.Println(movies)
		utils.WriteJSON(w, http.StatusOK, movies)
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
	}
}

// Получение ТОП фильмов (20) page = 1
func (h Handler) handlePopularFilms(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Инициализация переменной для получения из URL параметров
	query := r.URL.Query()
	// Получение page из URL
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

// Получение информации о фильме по ID
func (h Handler) handleMovieDetails(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получение ID из URL
	vars := mux.Vars(r)
	var id string = vars["id"]
	// Конвертация id в int
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

// Получение похожих фильмов
func (h Handler) handleMovieSimilar(w http.ResponseWriter, r *http.Request) {
	// Инициализация переменной для получения из URL параметров
	query := r.URL.Query()

	// Получение параметров из URL
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

	// Подготовка данных к отправке
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
