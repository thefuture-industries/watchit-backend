package user

import (
	"flicksfi/internal/interfaces"
	"flicksfi/internal/types"
	"flicksfi/internal/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	service        interfaces.IUser
	limiterService interfaces.ILimiter
}

func NewHandler(service interfaces.IUser, limiterService interfaces.ILimiter) *Handler {
	return &Handler{
		service:        service,
		limiterService: limiterService,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Вход в аккаунт
	router.HandleFunc("/user/check", h.handleCheckUser).Methods("POST")
	// регестрация аккаунт
	router.HandleFunc("/user/add", h.handleAddUser).Methods("POST")
	// обновление данных аккаунт
	router.HandleFunc("/user/update", h.handleUpdate).Methods("PUT")
}

// ----------------------
// ----------------------
// Проверка на наличие аккаунта
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
// Регестрация аккаунта
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

	// Шифрование секретного слова
	encryptSecretWord, err := utils.Encrypt(payload.SecretWord)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Получение пользователя
	user, _ := h.service.GetUserBySecretWord(encryptSecretWord)
	if user != nil {
		if err := h.limiterService.UpdateLimits(user.UUID); err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
			"uuid":     user.UUID,
			"username": user.UserName,
			"email":    user.Email,
		})
		return
	}

	// Проверка на существование пользователя по IP
	u, _ := h.service.GetUserByIP(payload.IPAddress)
	if u != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user already exists"))
		return
	}

	// Создание пользователя
	uuid := uuid.NewString()
	if err := h.service.CreateUser(types.User{
		UUID:       uuid,
		SecretWord: encryptSecretWord,
		IPAddress:  payload.IPAddress,
		Country:    payload.Country,
		RegionName: payload.RegionName,
		Zip:        payload.Zip,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	u, err = h.service.GetUserByUUID(uuid)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Отправка успешного выполнения
	utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"uuid":     uuid,
		"username": u.UserName,
		"email":    u.Email,
	})
}

// --------------------------------
// --------------------------------
// обновление данных аккаунт
// --------------------------------
func (h Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	// Проверка на кол-во запросов от пользователя
	limiter := utils.DDosPropperty()
	if limiter.Available() == 0 {
		utils.WriteError(w, http.StatusTooManyRequests, fmt.Errorf("too many requests"))
		return
	}

	// Получаем данные пользователя
	var payload *types.UserUpdate

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

	// Обновление данных пользователя
	if err := h.service.UserUpdate(*payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"uuid":     payload.UUID,
		"username": payload.Username,
		"email":    payload.Email,
	})
}
