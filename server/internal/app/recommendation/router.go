package recommendation

import (
	"flicksfi/internal/interfaces"
	"flicksfi/internal/types"
	"flicksfi/internal/utils"
	"flicksfi/pkg/movie"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	service     interfaces.IRecommendation
	userService interfaces.IUser
}

func NewHandler(service interfaces.IRecommendation, userService interfaces.IUser) *Handler {
	return &Handler{
		service:     service,
		userService: userService,
	}
}

func (h Handler) RegisterRoutes(router *mux.Router) {
	// Получение рекомендаций пользователя
	router.HandleFunc("/recommendations/{uuid}", h.handleGetRecommendations).Methods("GET")
	// Добавление рекомендаций пользователя
	router.HandleFunc("/recommendations", h.handleAddRecommendations).Methods("POST")
}

var total_page int = 500

// @Summary Get recommendations
// @Tags recommendation
// @Description Get recommendations user by UUID
// @ID get-recommendations
// @Accept json
// @Produce json
// @Param uuid path string true "UUID user"
// @Param page query string false "page"
// @Success 200 {object} []types.Movie
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /recommendations/{uuid} [get]
func (h Handler) handleGetRecommendations(w http.ResponseWriter, r *http.Request) {
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
	if page == "" {
		page = strconv.Itoa(rand.Intn(total_page) + 1)
	}

	rand.Seed(time.Now().UnixNano())

	// Получение сайта из URL
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// Получение рекомендаций пользователя
	recoms, err := h.service.GetRecommendation(uuid)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if len(recoms) == 0 {
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
		return
	}

	// Получение фильмов подходящих пользователю
	movies, err := h.service.GetMovieRecommendations(recoms)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, movies)
}

// @Summary Adding to recommendation
// @Tags recommendation
// @Description Adding to recommendation
// @ID add-recommendations
// @Accept json
// @Produce json
// @Param DTO body types.RecommendationAddPayload true "data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /recommendations [post]
func (h Handler) handleAddRecommendations(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.RecommendationAddPayload

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	// Получение текущего пользователя
	user, err := h.userService.GetUserByUUID(payload.UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if user == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	// Проверка на существования рекомендаций пользователя
	exists, err := h.service.IsRecommendation(payload.UUID, payload.Title)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	} else if exists {
		utils.WriteJSON(w, http.StatusNoContent, nil)
		return
	}

	// Запись рекомендаций пользователя
	if err := h.service.AddRecommendation(*payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Recommendation is created successfully"})
}
