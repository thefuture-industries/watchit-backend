// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package sync

import (
	"net/http"

	"go-user-service/internal/common/utils"
	"go-user-service/internal/packages"

	"github.com/noneandundefined/vision-go"
	"go.uber.org/zap"
)

type Handler struct {
	monitor *vision.Vision
	logger  *zap.Logger
	errors  *packages.Errors
}

func NewHandler(monitor *vision.Vision, logger *zap.Logger, errors *packages.Errors) *Handler {
	return &Handler{
		monitor: monitor,
		logger:  logger,
		errors:  errors,
	}
}

func (h Handler) SyncHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Syncing is active")
}
