package movie

import (
	"net/http"
	"watchit/httpx/pkg/httpx"
)

func (h *Handler) GetDetailsMovieHandler(w http.ResponseWriter, r *http.Request) error {
	httpx.HttpResponse(w, r, http.StatusOK, "MOVIEs")
	return nil
}
