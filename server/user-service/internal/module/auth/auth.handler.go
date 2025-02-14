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

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {

}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var payload *types.UserPayload

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
