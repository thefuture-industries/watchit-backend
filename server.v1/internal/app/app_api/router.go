package app_api

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
	service     interfaces.IAppAPI
	userService interfaces.IUser
}

func NewHandler(service interfaces.IAppAPI, userService interfaces.IUser) *Handler {
	return &Handler{
		service:     service,
		userService: userService,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Создания API KEY
	router.HandleFunc("/create/api_key", h.handleCreateAPIKEY).Methods("POST")
}

func (h *Handler) handleCreateAPIKEY(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получение данных от пользователя
	var payload *types.CreateAPIKEYPayload

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

	u, err := h.userService.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Создание api_key
	// api_key, err := h.service.CreateAPIKEY()
	// if err != nil {
	// 	utils.WriteError(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// Создание api_key
	var api_key string = h.service.CreateAPIKEY()

	// Запись в БД API_KEY
	if err := h.service.InsertAPIKEY(api_key, u.UUID); err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, api_key)
}
