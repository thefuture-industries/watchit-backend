package movie

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"watchit/httpx/pkg/httpx"
	"watchit/httpx/pkg/httpx/httperr"
)

func (h *Handler) GetDetailsMovieHandler(w http.ResponseWriter, r *http.Request) error {
	idParam := mux.Vars(r)["id"]
	if idParam == "" {
		return httperr.NotFound(fmt.Sprintf("movie "))
	}

	httpx.HttpResponse(w, r, http.StatusOK, "MOVIEs")
	return nil
}
