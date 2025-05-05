package auth

import (
	"go-user-service/internal/common/database"
	"go-user-service/internal/common/database/actions"
	"go-user-service/internal/common/utils"
	"go-user-service/internal/packages"
	"go-user-service/internal/types"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func (h Handler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	var payload *types.SigninPayload

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "Not all fields are filled in")
		return
	}

	// Получение пользователя
	queryStart := time.Now()
	isUsername, err := actions.GetUserByUsername(payload.Username)
	if err != nil {
		h.errors.HandleErrors(err, true)
		utils.WriteJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	h.monitor.VisionDBQuery(time.Since(queryStart))

	if isUsername == nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "The user was not found")
		return
	}

	pincode_hash, err := utils.Encrypt(payload.PINCODE)
	if err != nil {
		h.errors.HandleErrors(err, false)
		utils.WriteJSON(w, r, http.StatusInternalServerError, err)
		return
	}

	// Проверка на pincode
	if pincode_hash != isUsername.PINCODE {
		utils.WriteJSON(w, r, http.StatusBadRequest, "The user was not found")
		return
	}

	uuid_hash, err := utils.Encrypt(isUsername.UUID)
	if err != nil {
		h.errors.HandleErrors(err, false)
		utils.WriteJSON(w, r, http.StatusInternalServerError, err)
		return
	}

	// Создание и установка cookie
	cookie := &http.Cookie{
		Name:     "auth-token",
		Value:    uuid_hash,
		Path:     "/",
		Expires:  time.Now().Add(2 * 24 * time.Hour),
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)

	utils.WriteJSON(w, r, http.StatusOK, map[string]interface{}{
		"uuid":     isUsername.UUID,
		"username": isUsername.Username,
		"email":    isUsername.Email,
	})
}

func (h Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var payload *types.SignupPayload

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "Not all fields are filled in")
		return
	}

	queryStart := time.Now()
	isUsername, err := actions.GetUserByUsername(payload.Username)
	if err != nil {
		h.errors.HandleErrors(err, true)
		utils.WriteJSON(w, r, http.StatusInternalServerError, err)
		return
	}
	h.monitor.VisionDBQuery(time.Since(queryStart))

	if isUsername != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, "This username is busy, try another one")
		return
	}

	pincode_hash, err := utils.Encrypt(payload.PINCODE)
	if err != nil {
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
		return
	}

	user := database.Users{
		UUID:      uuid.New().String(),
		Username:  payload.Username,
		PINCODE:   pincode_hash,
		IPAddress: payload.IPAddress,
		Country:   payload.Country,
	}

	queryStart = time.Now()
	if err := actions.CreateUser(&user); err != nil {
		h.errors.HandleErrors(err, true)
		utils.WriteJSON(w, r, http.StatusBadRequest, err)
		return
	}
	h.monitor.VisionDBQuery(time.Since(queryStart))

	utils.WriteJSON(w, r, http.StatusCreated, "The user has been successfully created!")
}

func (h Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {}

func (h Handler) SignoutHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем куки с истекшим сроком действия
	cookie := &http.Cookie{
		Name:   "auth-token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	utils.WriteJSON(w, r, http.StatusOK, "You have successfully logged out.")
}
