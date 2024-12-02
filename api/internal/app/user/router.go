package user

import (
	"flick_finder/internal/interfaces"
	"flick_finder/internal/types"
	"flick_finder/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	service interfaces.IUser
}

func NewHandler(service interfaces.IUser) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Вход в аккаунт
	router.HandleFunc("/user/check", h.handleCheckUser).Methods("POST")
	// регестрация аккаунт
	router.HandleFunc("/user/add", h.handleAddUser).Methods("POST")
	// обновление данных аккаунт
	router.HandleFunc("/user/update", h.handleUpdate).Methods("PUT")
	// Созранения избранного
	router.HandleFunc("/favourites", h.handleFavouriteAdd).Methods("POST")
	// Получение избранных пользователя
	router.HandleFunc("/favourites", h.handleFavourites).Methods("GET")
}

// ----------------------
// ----------------------
// LOGIN ROUTER
// ----------------------
func (h *Handler) handleCheckUser(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.LoginUserPayload

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

	// Проверка
	u, err := h.service.CheckUser(*payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if u.ID == 0 {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}

	data := fmt.Sprintf("%s:%s", u.IPAddress, u.CreatedAt)

	utils.WriteJSON(w, http.StatusOK, data)
}

// ----------------------
// ----------------------
// REGISTRATION ROUTER
// ----------------------
func (h *Handler) handleAddUser(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.User

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

	user, _ := h.service.GetUserBySecretWord(payload.SecretWord)
	if user != nil {
		utils.WriteJSON(w, http.StatusCreated, user.UUID)
		return
	}

	// Создание пользователя
	uuid := uuid.NewString()
	if err := h.service.CreateUser(types.User{
		UUID:       uuid,
		SecretWord: payload.SecretWord,
		IPAddress:  payload.IPAddress,
		Lat:        payload.Lat,
		Lon:        payload.Lon,
		Country:    payload.Country,
		RegionName: payload.RegionName,
		Zip:        payload.Zip,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Отправка успешного выполнения
	utils.WriteJSON(w, http.StatusCreated, uuid)
}

// --------------------------------
// --------------------------------
// обновление данных аккаунт
// --------------------------------
func (h Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {

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
	user, err := h.service.GetUserByUUID(payload.UUID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// Запись в БД избранного
	if err := h.service.AddFavourite(*payload, user.UUID); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "success")
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

	// Получаем параметры запроса
	query := r.URL.Query()
	uuid := query.Get("uuid")

	favourites, err := h.service.Favourites(uuid)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, favourites)
}
