// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package admin

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h Handler) RegisterRoutes(router *mux.Router) {
	adminRouter := router.PathPrefix("/admin").Subrouter()

	adminRouter.HandleFunc("/vision/stats", h.MonitoringHandler).Methods(http.MethodGet)
}
