// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package sync

import (
	"net/http"

	"go-user-service/cmd/conf"
	"go-user-service/internal/common/utils"

	"go.uber.org/zap"
)

type Handler struct {
	monitor *conf.Vision
	logger  *zap.Logger
}

func NewHandler(monitor *conf.Vision, logger *zap.Logger) *Handler {
	return &Handler{
		monitor: monitor,
		logger:  logger,
	}
}

func (h Handler) SyncHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, "Syncing is active")
}
