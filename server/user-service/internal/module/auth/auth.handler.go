// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package auth

import (
	"fmt"
	"go-user-service/internal/common/database"
	"go-user-service/internal/common/database/actions"
	"go-user-service/internal/common/types"
	"go-user-service/internal/common/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
)

type Handler struct {
	monitor *vision.Vision
	logger  *zap.Logger
}

func NewHandler(monitor *vision.Vision, logger *zap.Logger) *Handler {
	return &Handler{
		monitor: monitor,
		logger:  logger,
	}
}

func (h Handler) SigninHandler(w http.ResponseWriter, r *http.Request) {
	var payload *types.SigninPayload

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("%v", errors))
		return
	}

	// Получение пользователя
	queryStart := time.Now()
	isUsername, err := actions.GetUserByUsername(payload.Username)
	if err != nil {
		ctx := utils.SetErrorInContext(r.Context(), err)
		r = r.WithContext(ctx)
		h.monitor.VisionError(err)
		// h.monitor.VisionDBError()
		// h.logger.Error("[DB ERROR]", zap.Error(err))
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	h.monitor.VisionDBQuery(time.Since(queryStart))

	if isUsername == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("the user was not found"))
		return
	}

	pincode_hash, err := utils.Encrypt(payload.PINCODE)
	if err != nil {
		h.logger.Error("[HASH ERROR]", zap.Error(err))
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// Проверка на pincode
	if pincode_hash != isUsername.PINCODE {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("the user was not found"))
		return
	}

	uuid_hash, err := utils.Encrypt(isUsername.UUID)
	if err != nil {
		h.logger.Error("[HASH ERROR]", zap.Error(err))
		utils.WriteError(w, http.StatusInternalServerError, err)
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

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"uuid":     isUsername.UUID,
		"username": isUsername.Username,
		"email":    isUsername.Email,
	})
}

func (h Handler) SignupHandler(w http.ResponseWriter, r *http.Request) {
	var payload *types.SignupPayload

	// Отправляем пользователю ошибку, что не все поля заполнены
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Валидация данных от пользователя
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("%v", errors))
		return
	}

	isUsername, err := actions.GetUserByUsername(payload.Username)
	if err != nil {
		h.monitor.VisionDBError()
		h.logger.Error("[DB ERROR]", zap.Error(err))
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if isUsername != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("this username is busy, try another one"))
		return
	}

	pincode_hash, err := utils.Encrypt(payload.PINCODE)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := database.Users{
		UUID:      uuid.New().String(),
		Username:  payload.Username,
		PINCODE:   pincode_hash,
		IPAddress: payload.IPAddress,
		Country:   payload.Country,
	}

	if err := actions.CreateUser(&user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, "the user has been successfully created!")
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

	utils.WriteJSON(w, http.StatusOK, "You have successfully logged out.")
}
