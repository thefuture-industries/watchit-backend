// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package sync

import (
	"net/http"

	"go-user-service/internal/common/utils"

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

func (h Handler) SyncHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Syncing is active")
}
