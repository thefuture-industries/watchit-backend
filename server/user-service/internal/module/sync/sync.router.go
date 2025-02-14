// *---------------------------------------------------------------------------------------------
//  *  Copyright (c). All rights reserved.
//  *  Licensed under the MIT License. See License.txt in the project root for license information.
//  *--------------------------------------------------------------------------------------------*

package sync

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/sync", SyncHandler).Methods(http.MethodGet)
}
