package favourite

import (
	"flicksfi/internal/interfaces"
	"flicksfi/internal/types"
	"flicksfi/internal/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	service               interfaces.IFavourite
	userService           interfaces.IUser
	recommendationService interfaces.IRecommendation
}

func NewHandler(service interfaces.IFavourite, userService interfaces.IUser, recommendationService interfaces.IRecommendation) *Handler {
	return &Handler{
		service:               service,
		userService:           userService,
		recommendationService: recommendationService,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Созранения избранного
	router.HandleFunc("/favourites", h.handleFavouriteAdd).Methods("POST")
	// Получение избранных пользователя
	router.HandleFunc("/favourites/{uuid}", h.handleFavourites).Methods("GET")
	// Удаление избранных пользователя
	router.HandleFunc("/favourites", h.handleFavouriteDelete).Methods("DELETE")
}

// --------------------------------
// --------------------------------
// Созранения избранного
// --------------------------------
func (h Handler) handleFavouriteAdd(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.FavouriteAddPayload

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

	// Получение текущего пользователя
	user, err := h.userService.GetUserByUUID(payload.UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Проверка на то что избранного нет у uuid
	if err := h.service.CheckFavourites(user.UUID, payload.MovieID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Запись в БД избранного
	if err := h.service.AddFavourite(*payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Favourite added successfully"})
}

// --------------------------------
// --------------------------------
// Получение избранных фильмов пользователя
// --------------------------------
func (h Handler) handleFavourites(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получение сайта из URL
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	favourites, err := h.service.Favourites(uuid)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, favourites)
}

// ---------------------------------------
// ---------------------------------------
// Удаление избранных фильмов пользователя
// ---------------------------------------
func (h Handler) handleFavouriteDelete(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.FavouriteDeletePayload

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

	// Получение пользователя
	user, _ := h.userService.GetUserByUUID(payload.UUID)
	if user == nil {
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("user not found"))
		return
	}

	// Удаление избранного
	if err := h.service.DeleteFavourite(*payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Favourite deleted successfully"})
}
